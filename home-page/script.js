let videoContainer = document.getElementById('videoContainer');
let sidebarButtons = document.querySelectorAll('.sidebar button');
let peer;
let localStream;
let roomIdInput = document.getElementById('roomId');
let createRoomButton = document.getElementById('createRoom');
let joinRoomButton = document.getElementById('joinRoom');
let currentCall = null;
let isCameraOn = true; 
let isMicOn = true; 

videoContainer.style.left = '53%';
videoContainer.style.transform = 'translateX(-53%)';

function toggleVideoPosition() {
    if (videoContainer.style.left === '53%') {
      videoContainer.style.left = '80%';
      videoContainer.style.transform = 'translateX(0)';
    } else {
      videoContainer.style.left = '53%';
      videoContainer.style.transform = 'translateX(-53%)';
    }
  }
  
function toggleVideoAndEditor(event) {
  const codeEditor = document.getElementById('codeEditor');
  if (event.target === openCodeEditor) {
    
    if (codeEditor.style.display === 'none' || !codeEditor.style.display) {
      codeEditor.style.display = 'block';   
    } else {
      codeEditor.style.display = 'none';   
    }
  } else if (event.target === videoContainer) {
   
    videoContainer.classList.toggle('moved-position'); 
  }
}

const openCodeEditor = document.getElementById('openCodeEditor'); 
const viewGithub = document.getElementById('viewGithub'); 

openCodeEditor.addEventListener('click', toggleVideoPosition);
viewGithub.addEventListener('click', toggleVideoPosition);
openCodeEditor.addEventListener('click', toggleVideoAndEditor);
viewGithub.addEventListener('click', toggleVideoAndEditor);


videoContainer.addEventListener('click', toggleVideoPosition);


function initializePeer(isCreatingRoom) {
  if (peer) {
    peer.destroy();
  }

  peer = new Peer(undefined, {
    config: {
      'iceServers': [
        { urls: 'stun:stun.l.google.com:19302' },
        { urls: 'stun:stun1.l.google.com:19302' },
        { urls: 'stun:stun2.l.google.com:19302' },
        { urls: 'stun:stun3.l.google.com:19302' },
        { urls: 'stun:stun4.l.google.com:19302' }
      ]
    },
    debug: 3
  });

  peer.on('open', (id) => {
    console.log('My peer ID is: ' + id);
    if (isCreatingRoom) {
      roomIdInput.value = id;
      document.getElementById('createRoomStatus').innerText = `Room Created: ${id}`;
    }
  });

  peer.on('call', (call) => {
    currentCall = call;
    call.answer(localStream);
    handleCall(call);
  });

  peer.on('error', (error) => {
    console.error('PeerJS error:', error);
    alert(`Connection error: ${error.type}`);
  });

  peer.on('disconnected', () => {
    console.log('Connection lost. Please reconnect');
    document.getElementById('createRoomStatus').innerText = 'Connection lost. Please refresh the page.';
  });
}


function handleCall(call) {
  call.on('stream', (remoteStream) => {
    const remoteVideo = document.getElementById('remoteVideo');
    remoteVideo.srcObject = remoteStream;
  });

  call.on('close', () => {
    const remoteVideo = document.getElementById('remoteVideo');
    remoteVideo.srcObject = null;
  });

  call.on('error', (error) => {
    console.error('Call error:', error);
    alert(`Call error: ${error}`);
  });
}


async function startVideo() {
  try {
    localStream = await navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true
    });
    const localVideo = document.getElementById('localVideo');
    localVideo.srcObject = localStream;
    localVideo.play().catch(e => console.error('Error playing local video:', e));
    return true;
  } catch (error) {
    console.error('Error accessing media devices:', error);
    alert('Error accessing camera/microphone: ' + error.message);
    return false;
  }
}


createRoomButton.addEventListener('click', async () => {
  const videoStarted = await startVideo();
  if (!videoStarted) return;

  initializePeer(true);
});


joinRoomButton.addEventListener('click', async () => {
  if (!roomIdInput.value) {
    alert('Please enter a Room ID!');
    return;
  }

  const videoStarted = await startVideo();
  if (!videoStarted) return;

  initializePeer(false);

  peer.on('open', () => {
    const call = peer.call(roomIdInput.value, localStream);
    if (call) {
      currentCall = call;
      handleCall(call);
      document.getElementById('createRoomStatus').innerText = `Joined Room: ${roomIdInput.value}`;
    } else {
      alert('Failed to connect to peer. Please check the Room ID and try again.');
    }
  });
});


function endCall() {
  if (currentCall) {
    currentCall.close();
    currentCall = null;
    document.getElementById('createRoomStatus').innerText = 'Call ended';
  }

  
  if (localStream) {
    localStream.getTracks().forEach(track => track.stop());
    const localVideo = document.getElementById('localVideo');
    localVideo.srcObject = null;
  }
}


function toggleCamera() {
  isCameraOn = !isCameraOn;

  if (isCameraOn) {
    const constraints = { video: true, audio: true };
    navigator.mediaDevices.getUserMedia(constraints).then((stream) => {
      localStream.getVideoTracks()[0].enabled = true;
      document.getElementById('localVideo').srcObject = localStream;
    }).catch((error) => {
      console.error('Error enabling camera:', error);
      alert('Failed to enable camera');
    });
  } else {
    localStream.getVideoTracks()[0].enabled = false;
    document.getElementById('localVideo').srcObject = localStream;
  }
}


function toggleMic() {
  isMicOn = !isMicOn;

  if (isMicOn) {
    localStream.getAudioTracks().forEach(track => track.enabled = true);
  } else {
    localStream.getAudioTracks().forEach(track => track.enabled = false);
  }
}


document.getElementById('endCallButton').addEventListener('click', endCall);
document.getElementById('toggleCameraButton').addEventListener('click', toggleCamera);
document.getElementById('toggleMicButton').addEventListener('click', toggleMic);


window.onbeforeunload = () => {
  if (localStream) {
    localStream.getTracks().forEach(track => track.stop());
  }
  if (currentCall) {
    currentCall.close();
  }
  if (peer) {
    peer.destroy();
  }
};

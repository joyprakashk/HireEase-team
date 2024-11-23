import streamlit as st
import datetime
import os
import matplotlib.pyplot as plt
from groq import Groq
import google.generativeai as genai
from PIL import Image  

API_KEY_GROQ = "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"
API_KEY_GENAI = "AIzaSyC4f-d-Igv6UWdHKoMgZcNfRTeQBFVgtUw"

client = Groq(api_key=API_KEY_GROQ)
genai.configure(api_key=API_KEY_GENAI)

st.set_page_config(page_title="NutriGuide.AI", page_icon="🍎", layout="wide")

app_mode = st.sidebar.selectbox(
    "Choose a feature", 
    [
        "Home",
        "Personalized Meal Plan",
        "Food Nutrient and Calorie Estimator",
        "Food Allergen Checker",
        "Diet Recommendation for Chronic Conditions",
        "Personalized Diet and Fitness Syncing"
    ]
)

def load_dynamic_css():
    st.markdown(
        """
        <style>
        /* General Dark Mode Setup */
        .stApp {
            background-color: #2c3e50;  /* Dark background */
        }

        /* Sidebar Styling */
        .css-1d391kg {
            background-color: #34495e;  /* Darker sidebar */
        }


        /* Centering Title */
        .main-bg h1 {
            color: #ffffff;  /* White text color for title */
            text-align: center;
            font-size: 3rem;
            font-weight: 600;
        }

        /* Button Animation */
        .stButton>button {
            background-color: #16a085;
            color: #2c3e50;
            border-radius: 5px;
            transition: transform 0.3s ease-in-out;
        }
        
        .stButton>button:hover {
            transform: scale(1.05);
        }

        /* List Styling and Hover Effects */
        ul {
            margin: 20px 0;
            padding-left: 30px;
        }
        
        ul li {
            color: #bdc3c7;
            font-size: 1.1rem;
            padding: 5px 0;
            transition: color 0.3s ease;
        }

        ul li:hover {
            color: #1abc9c;  /* Green color on hover */
        }

        /* About Us Section Styling */
        .feature-bg {
            background-color: #34495e;
            padding: 25px;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            margin-bottom: 30px;
        }

        .feature-bg h2 {
            color: #1abc9c;
            font-size: 1.8rem;
            margin-bottom: 10px;
        }

        .feature-bg p {
            color: #bdc3c7;
            font-size: 1.1rem;
        }

        /* Section Styling */
        .feature-section {
            padding: 30px;
            margin-top: 40px;
            border-radius: 12px;
            background-color: #34495e;
            box-shadow: 0 5px 10px rgba(0, 0, 0, 0.2);
        }
        .feature-section ul {
            padding-left: 20px;
        }

        /* Footer Styling */
        .footer {
            text-align: center;
            padding: 20px;
            margin-top: 50px;
            background-color: #34495e;
            border-radius: 10px;
            color: #ecf0f1;
        }
        </style>
        """,
        unsafe_allow_html=True
    )
load_dynamic_css()

if app_mode == "Home":
    # Display full-width image
    st.image("Green Illustrative Food Journal Presentation.jpg", use_column_width=True)

    # Below the hero image, you can add other sections
    st.markdown('<div class="main-bg">', unsafe_allow_html=True)

    st.markdown('<h1>Welcome to NutriGuide.AI</h1>', unsafe_allow_html=True)
    st.markdown(
        """
        <div class="feature-bg">
            <h2>Your Personalized Path to Better Nutrition</h2>
            <p>
                NutriGuide.AI is your ultimate personal health and diet assistant powered by AI. Our advanced algorithms provide personalized solutions for your health and nutrition needs. With NutriGuide.AI, you get accurate recommendations to help you achieve a healthier, happier lifestyle!
            </p>
            <p>
                Here’s what we offer:
            </p>
            <ul>
                <li><b>Personalized Meal Plans:</b> Tailored to your dietary preferences, goals, and health requirements.</li>
                <li><b>Food Calorie Estimation:</b> Instantly estimate the calories in your meals, helping you track your daily intake and stay on top of your nutrition.</li>
                <li><b>Allergen Checker:</b> Our AI-powered allergen checker helps you avoid harmful ingredients and allergens in your food.</li>
                <li><b>Chronic Conditions Support:</b> Tailored diet recommendations based on your specific health conditions, such as diabetes, heart disease, and more.</li>
                <li><b>Fitness Syncing:</b> Sync your fitness routine with your meal plan, ensuring that your nutrition supports your fitness goals.</li>
            </ul>
        </div>

        <div class="feature-bg">
            <h2>About Us</h2>
            <p>
                NutriGuide.AI was founded with the mission to make healthy eating easier, smarter, and personalized for everyone. Our team of experts in nutrition, artificial intelligence, and health technology have come together to create an app that helps you make informed food choices, track your meals, and stay on top of your health goals.
            </p>
        </div>

        <div class="feature-section">
            <h2>Why Choose NutriGuide.AI?</h2>
            <ul>
                <li><b>Personalization at its Core:</b> Every recommendation is tailored to your preferences and health needs, ensuring you get the best advice.</li>
                <li><b>AI-powered Insights:</b> Our advanced AI-driven algorithms analyze your dietary habits, providing data-driven solutions to optimize your health.</li>
                <li><b>Comprehensive Approach:</b> We offer features that go beyond simple meal tracking, including fitness syncing, allergen checking, and chronic conditions support.</li>
                <li><b>User-friendly Interface:</b> NutriGuide.AI has a sleek, intuitive design that ensures easy navigation and effortless use for everyone, from beginners to experts.</li>
            </ul>
        </div>

        <div class="feature-section">
            <h2>Features</h2>
            <ul>
                <li><b>Personalized Meal Plans</b> – Tailored to your dietary preferences, lifestyle, and health goals.</li>
                <li><b>Food Calorie Estimator</b> – Quickly calculate the nutritional value of your meals.</li>
                <li><b>Food Allergen Checker</b> – Scan your food and check for potential allergens.</li>
                <li><b>Chronic Conditions Support</b> – Get diet recommendations for specific conditions like diabetes, hypertension, and more.</li>
                <li><b>Fitness Syncing</b> – Sync with your fitness routine to optimize diet and exercise together.</li>
            </ul>
        </div>

        <div class="feature-bg">
            <h2>Ready to Start Your Nutrition Journey?</h2>
            <p>Join the thousands of users already using NutriGuide.AI to take control of their health. Try it today!</p>
            <div class="stButton"><button>Get Started</button></div>
        </div>
    """,
    unsafe_allow_html=True
)

    # Footer Section
    st.markdown(
        """
        <div class="footer">
            <p>🍎 NutriGuide.AI © 2024 | All Rights Reserved</p>
        </div>
        """,
        unsafe_allow_html=True
    )

# Feature 1: Personalized Meal Plan
if app_mode == "Personalized Meal Plan":
    st.markdown('<div class="feature-bg"><h2>Personalized Meal Plan</h2></div>', unsafe_allow_html=True)

    
    age = st.number_input("Enter your age:", min_value=0, max_value=120)
    sex = st.selectbox("Select your sex:", ["Male", "Female", "Other"])
    weight_goal = st.selectbox("Weight goal:", ["Gain", "Loss"])
    eating_habits = st.selectbox("Eating habits:", ["Omnivore", "Flexitarian", "Pescatarian", "Vegetarian", "Vegan"])
    dietary_restrictions = st.selectbox("Dietary restrictions:", ["None", "Mediterranean", "Paleo", "Whole30", "Low Carb", "High Carb", "Gluten-Free", "Lactose-Free", "Raw Food", "Alkaline Food"])
    intermittent_fast_timing = st.selectbox("Select your intermittent fasting schedule:", ["None", "16/8", "18/6", "20/4"])
    traditional_diet = st.selectbox("Choose traditional diet:", ["None", "Ayurvedic", "Macrobiotic", "Halal", "Kosher"])
    duration = st.selectbox("Meal plan duration:", ["One day", "One week", "One month"])

    if st.button("Get Meal Plan", key="meal_plan_btn"):
        try:
            messages = [
                {
                    "role": "user",
                    "content": f"Write a personalized meal generation for the following parameters\n"
                               f"1. Age: {age}\n"
                               f"2. Sex: {sex}\n"
                               f"3. Weight goal: {weight_goal}\n"
                               f"4. Food eating habits: {eating_habits}\n"
                               f"5. Dietary restrictions: {dietary_restrictions}\n"
                               f"6. Intermittent fasting :{intermittent_fast_timing}\n"
                               f"7. Traditional diet: {traditional_diet}\n"
                               f"8. duration :{duration}\n"
                },
                {
                    "role": "assistant",
                    "content":"dont inlucde to consult wiht dietictian message ,start with your personlized meal plan is\n"
                              "Provide a meal plan based on the above parametersand also include the cost of the meals in INR\n"
                              "make sure first the dish must be mentioned followed by ingeredints and its cost then approximation of entire dish\n"
                              "when its intermittent fasting make sure to take time which user provides\n"
                              "provide deatils of the paramter choosen by the user before giving the diet \n"
                              "If i ask in days give me in the complete day. i want youto give me all the data which asked "

                }
            ]
            completion = client.chat.completions.create(
                model="llama-3.2-90b-text-preview",
                messages=messages,
                temperature=1,
                max_tokens=1024
            )
            meal_plan = completion.choices[0].message.content
            st.write(meal_plan)
        except Exception as e:
            st.error(f"Error: {e}")


# Feature 2: Food Nutrient and Calorie Estimator
if app_mode == "Food Nutrient and Calorie Estimator":
    st.markdown('<div class="feature-bg"><h2>Food Nutrient and Calorie Estimator</h2></div>', unsafe_allow_html=True)

    st.write("""
        Upload an image of your food, and our AI will analyze the nutrient composition, including calories, protein, fats, carbohydrates, vitamins, and minerals. 
        Additionally, you'll receive a meal balance assessment.
    """)

    image_file = st.file_uploader("Upload an image of food:", type=["jpg", "jpeg", "png"])

    if image_file:
        # Ensure the 'temp' directory exists
        temp_dir = "temp"
        if not os.path.exists(temp_dir):
            os.makedirs(temp_dir)

        temp_image_path = os.path.join(temp_dir, image_file.name)
        with open(temp_image_path, "wb") as f:
            f.write(image_file.getbuffer())

        st.image(image_file, caption="Uploaded Image", use_column_width=True)

        if st.button("Analyze Nutrients"):
            try:
                
                myfile = genai.upload_file(temp_image_path)
                model = genai.GenerativeModel("gemini-1.5-flash")

                result = model.generate_content(
                    [
                        myfile,
                        "\n\n",
                        """You are an expert nutritionist. Analyze the food dish or dishes in this image and provide a detailed breakdown of the nutrient composition, and you do not need to provide exact composition; you can give expected ranges including expected values for calories, protein, fats, carbohydrates, vitamins, and minerals. 
                        Provide the analysis in the following format in bullet points:
                        - Food Item: 
                          - Calories: X
                          - Protein: Yg
                          - Fats: Zg
                          - Carbohydrates: Ag
                          - Vitamins: [List]
                          - Minerals: [List]
                        """
                    ]
                )

                st.success("Nutrient analysis completed!")
                st.write(result.text)

                balance_result = model.generate_content(
                    [
                        result.text,
                        "\n\n",
                        """Based on the nutrient analysis provided, categorize the meal as one of the following:
                        - Not Healthy
                        - Healthy
                        - Nutritious
                        - Somewhat balanced
                        - Highly balanced and nutritious
                        Provide a short explanation for your assessment and, if required, give tips to make it more balanced.
                        """
                    ]
                )

                st.markdown("<div style='background-color: #f0f0f0; padding: 10px; border-radius: 5px;'>"
                            "<strong>Meal Balance Assessment:</strong><br>"
                            f"{balance_result.text}"
                            "</div>", unsafe_allow_html=True)

            except Exception as e:
                st.error(f"An error occurred: {e}")

        if os.path.exists(temp_image_path):
            os.remove(temp_image_path)


# Feature 3: Food Allergen Checker
elif app_mode == "Food Allergen Checker":
    st.markdown('<div class="feature-bg"><h2>Food Allergen Checker</h2></div>', unsafe_allow_html=True)
    food_item = st.text_input("Enter the food dish you want to check (e.g., Pizza):")
    allergen = st.text_input("Enter any specific allergens if you have (optional):")
    image_file = st.file_uploader("Upload an image of the food dish (optional):", type=["jpg", "jpeg", "png"])

    if st.button("Check Allergens"):
        temp_image_path = None  # Initialize variable for temporary image path
        if food_item or image_file:
            prompt = ""
            if allergen:
                if food_item:
                    prompt += f"The dish '{food_item}' is being analyzed. I am allergic to '{allergen}'. Does this dish contain my allergen? Please provide a simple 'yes', 'no', 'high probability', or 'low chances', and 4-5 lines of explanation with a warning about avoiding '{allergen}' in the dish."
                if image_file:
                    temp_image_path = os.path.join("temp", image_file.name)
                    with open(temp_image_path, "wb") as f:
                        f.write(image_file.getbuffer())
                    prompt += f"\nThis is the food '{temp_image_path}' I am going to eat. I am allergic to '{allergen}'. Just tell if the dish can generally contain '{allergen}'. Give a one-word answer followed by a detailed explanation."
            else:
                if food_item:
                    prompt += f"Analyze the dish '{food_item}' for potential allergens. What ingredients might cause allergic reactions? Please provide a summary."
                if image_file:
                    temp_image_path = os.path.join("temp", image_file.name)
                    with open(temp_image_path, "wb") as f:
                        f.write(image_file.getbuffer())
                    prompt += f"\nThis is the food '{temp_image_path}' I am going to eat. What ingredients might cause allergic reactions? Please provide a summary."

            # Execute AI model
            if prompt:
                model = genai.GenerativeModel("gemini-1.5-flash")
                try:
                    if temp_image_path:
                        myfile = genai.upload_file(temp_image_path)
                        result = model.generate_content([myfile, "\n\n", prompt])
                    else:
                        result = model.generate_content(prompt)
                    
                    response_lines = result.text.splitlines()
                    simple_response = response_lines[0] if response_lines else "No response received."
                    explanation = "\n".join(response_lines[1:]) if len(response_lines) > 1 else ""

                    # Display responses in colored boxes
                    if "yes" in simple_response.lower():
                        response_color = "#FF4D4D"
                    elif "high probability" in simple_response.lower():
                        response_color = "#FF4D4D"
                    elif "no" in simple_response.lower():
                        response_color = "#ADD8E6"
                    elif "low chances" in simple_response.lower():
                        response_color = "#B0E0E6"
                    else:
                        response_color = "#D3D3D3"

                    st.success("Allergen check completed!")
                    st.markdown(
                        f"<div style='border: 2px solid {response_color}; padding: 10px; border-radius: 5px; background-color: {response_color};'>Response: {simple_response}</div>",
                        unsafe_allow_html=True
                    )
                    if explanation:
                        st.markdown(
                            f"<div style='border: 2px solid #2196F3; padding: 10px; border-radius: 5px;'><strong>Explanation:</strong><br>{explanation}</div>",
                            unsafe_allow_html=True
                        )
                except Exception as e:
                    st.error(f"An error occurred during the allergen check: {e}")
                finally:
                    if temp_image_path and os.path.exists(temp_image_path):
                        os.remove(temp_image_path)
            else:
                st.error("Unable to generate a prompt for the AI. Please check your inputs.")
        else:
            st.error("Please enter a food dish or upload a photo to check.")

# Feature 4: Diet Recommendation for Chronic Conditions
elif app_mode == "Diet Recommendation for Chronic Conditions":
    st.markdown('<div class="feature-bg"><h2>Diet Recommendation for Chronic Conditions</h2></div>', unsafe_allow_html=True)

    condition = st.text_input("Enter your chronic condition (e.g., Diabetes, Hypertension):")
    severity = st.selectbox("Select the severity level:", ["Mild", "Moderate", "Severe"])
    goals = st.text_input("Health goals related to this condition (e.g., lower blood pressure, manage blood sugar)")
    food = st.selectbox("What type of meal do you like?", ["Veg", "Non-Veg", "Vegan"])

    if st.button("Get Diet Recommendation"):
        try:
            messages = [
                {
                    "role": "user",
                    "content": (
                        f"Provide detailed information for {condition},{severity},{goals}{food}.\n"
                        "1. Conditional-specific nutritional guidelines\n"
                        "2. Customized meal planning. Include options for day, week, or month.\n"
                        "3. Symptom management tips\n"
                        "4. Educational content and lifestyle tips.\n"
                        "Please format the response accordingly."
            )
                }
            ]
            completion = client.chat.completions.create(
                model="llama-3.2-90b-text-preview",
                messages=messages,
                temperature=1,
                max_tokens=1024
            )
            diet_recommendation = completion.choices[0].message.content
            st.write(diet_recommendation)
        except Exception as e:
            st.error(f"Error: {e}")

# Feature 5: Personalized Diet and Fitness Syncing
elif app_mode == "Personalized Diet and Fitness Syncing":
    st.markdown('<div class="feature-bg"><h2>Personalized Diet and Fitness Syncing</h2></div>', unsafe_allow_html=True)
    # User Inputs
    age = st.number_input("Age", min_value=1, max_value=100, value=30)
    gender = st.selectbox("Gender", ["Male", "Female", "Other"])
    weight = st.number_input("Weight (kg)", min_value=20, max_value=200, value=70)
    height = st.number_input("Height (cm)", min_value=100, max_value=250, value=165)
    body_fat_percentage = st.number_input("Body Fat Percentage (%)", min_value=1, max_value=50, value=20)
    sleep_hours = st.number_input("Average Sleep Hours per Night", min_value=1, max_value=24, value=7)
    stress_level = st.selectbox("Stress Level", ["Low", "Medium", "High"])
    
    goal = st.selectbox("Diet Goal", ["Weight Loss", "Weight Gain", "Maintenance"])
    activity_level = st.selectbox("Activity Level", ["Sedentary", "Lightly Active", "Moderately Active", "Very Active"])
    
    protein_goal = st.number_input("Daily Protein Goal (g)", min_value=10, max_value=300, value=100)
    carb_goal = st.number_input("Daily Carbohydrate Goal (g)", min_value=20, max_value=500, value=250)
    fat_goal = st.number_input("Daily Fat Goal (g)", min_value=10, max_value=200, value=70)
    
    dietary_preferences = st.selectbox("Dietary Preferences", ["No Preference", "Vegetarian", "Vegan", "Low-Carb", "High-Protein"])
    dietary_restrictions = st.multiselect("Dietary Restrictions", ["Gluten-Free", "Dairy-Free", "Nut-Free", "Shellfish-Free"])
  
    time_frame_type = st.selectbox("Select Time Frame Type", ["Days", "Weeks", "Months"])
    if time_frame_type == "Days":
        time_frame = st.number_input("Time Frame (in days)", min_value=1, max_value=365, value=7)
    elif time_frame_type == "Weeks":
        time_frame = st.number_input("Time Frame (in weeks)", min_value=1, max_value=52, value=1) * 7  
    else:
        time_frame = st.number_input("Time Frame (in months)", min_value=1, max_value=12, value=1) * 30 

    calorie_goal = st.number_input("Daily Calorie Goal (kcal)", min_value=1000, max_value=5000, value=1800)
    
    recent_workout_data = st.text_area("Enter recent workout data (e.g., duration, type, intensity):", "")
   
    hydration_goal = st.slider("Daily Hydration Goal (L)", min_value=1.0, max_value=5.0, value=2.5)
   
    tracking_period = st.selectbox("Choose Tracking Period", ["Daily", "Weekly"])  # Ensure this is defined before being used in prompt
    
    st.sidebar.header("Daily Motivational Tip")
    motivational_tips = [
        "Stay consistent! Small progress adds up.",
        "Remember to rest and recover.",
        "Hydrate! Water is key to wellness.",
        "Take it one day at a time, you've got this!",
        "Celebrate small victories!",
    ]
    st.sidebar.write(motivational_tips[datetime.datetime.now().day % len(motivational_tips)])
    
    if st.button("Generate Diet and Fitness Plan"):
        prompt = f"""
        You are a nutrition and fitness expert. Based on the following information, create a detailed personalized diet and fitness plan:
        - Age: {age}
        - Gender: {gender}
        - Weight: {weight} kg
        - Height: {height} cm
        - Body Fat Percentage: {body_fat_percentage}%
        - Sleep: {sleep_hours} hours per night
        - Stress Level: {stress_level}
        - Goal: {goal}
        - Activity Level: {activity_level}
        - Dietary Preferences: {dietary_preferences}
        - Dietary Restrictions: {", ".join(dietary_restrictions)}
        - Time Frame: {time_frame} days
        - Daily Calorie Goal: {calorie_goal} kcal
        - Protein Goal: {protein_goal} g
        - Carbohydrate Goal: {carb_goal} g
        - Fat Goal: {fat_goal} g
        - Recent Workout Data: {recent_workout_data}
        - Daily Hydration Goal: {hydration_goal} L
        
        For each day of the time frame, provide:
        1. Detailed dietary recommendations, considering recent workout data and dietary restrictions.
        2. Breakdown of meals: breakfast, lunch, dinner, and snacks, with specific foods listed.
        3. Calorie and nutrient breakdown for each meal (e.g., protein, carbs, fats).
        4. Hydration goals and reminders.
        5. A summary of daily and weekly intake to help track calories and nutrient balance, including insights for {tracking_period} tracking.
        6. Provide tips on food choices and alternatives to reach the calorie goal.
        7. Suggest food alternatives based on available ingredients and preferences.
        8. Include a brief motivational tip for each day to help the user stay on track.
        """
        with st.spinner("Generating your personalized diet and fitness plan..."):
            try:
                # Use Groq client to generate the plan
                messages = [{"role": "user", "content": prompt}]
                completion = client.chat.completions.create(
                    model="llama-3.2-90b-text-preview",
                    messages=messages,
                    temperature=1,
                    max_tokens=1024
                )
                plan = completion.choices[0].message.content
                st.success("Diet and fitness plan generated successfully!")
                st.write(plan)
            except Exception as e:
                st.error(f"An error occurred: {e}")
    
    if tracking_period == "Weekly":
        st.header("Progress Overview")
        days = [f"Day {i}" for i in range(1, time_frame + 1)]
        calorie_intake = [calorie_goal + (i % 5 * 10 - 20) for i in range(time_frame)] 
        fig, ax = plt.subplots(figsize=(12, 6))
        ax.plot(days, calorie_intake, label="Calorie Intake", color='blue', marker='o')
        ax.set_xlabel("Days")
        ax.set_ylabel("Calories (kcal)")
        ax.set_title("Weekly Calorie Intake")
        ax.legend()
        st.pyplot(fig)

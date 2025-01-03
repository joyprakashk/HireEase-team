import pandas as pd
import google.generativeai as genai
import os
import fitz  

genai.configure(api_key="AIzaSyC4f-d-Igv6UWdHKoMgZcNfRTeQBFVgtUw")
model = genai.GenerativeModel("gemini-1.5-flash")

def recommend_companies(data, user_skills, skill_match_threshold=0.5):
    user_skills = [skill.strip().lower() for skill in user_skills]

    total_skills = len(user_skills)
    
    recommendations = data[ 
        data['Skills Required'].apply(
            lambda x: (
                sum(skill in x.lower() for skill in user_skills) / total_skills
            ) >= skill_match_threshold if pd.notnull(x) else False
        )
    ]
    return recommendations

def format_detailed_recommendations(recommendations, user_skills):
    detailed_recommendations = []
    for index, row in recommendations.iterrows():
        matched_skills = [
            skill for skill in user_skills if skill.lower() in row['Skills Required'].lower()
        ]
        match_count = len(matched_skills)
        total_skills_required = len(row['Skills Required'].split(', '))

        duration = row.get('Duration', 'Not specified')
        stipend = row.get('Stipend', 'Not specified')
        
        detailed_recommendations.append(f"""
{index+1}. **{row['Profile']}**  
{row['Company']} · {row['Location']}  
Duration: {duration} | Stipend: {stipend}  

**Skills Match:** {match_count} of {total_skills_required} skills match your profile.  
**Skills Required:** {row['Skills Required']}  

---
""")
    return "\n".join(detailed_recommendations)

def extract_text_from_docx(docx_file_path):
    doc = fitz.open(docx_file_path)  
    text = ""
    for page in doc:
        text += page.get_text("text")  
    return text

def compare_resume_to_job(resume_file_content, company_skills):
    prompt = f"""
     Please evaluate the following documents for skills alignment and determine a "Pass" or "Fail" result based on the outlined criteria:

**Document 1: Resume**
{resume_file_content}

**Document 2: Job Requirements**
{company_skills}

### Evaluation Criteria:

2. **Missing Skills (if Fail):** Identify only the critical skills from the job requirements that are missing in the resume.
3. **Recommendation (if Fail):** Suggest one or two brief, actionable improvements to align the resume better with the job requirements.
1. **Skills Match:** 
   - If **70% or more** of the skills from the job requirements are present in the resume, return "Pass."
   - If the resume includes **all of the following key skills**, automatically return "Pass":
     - Programming Languages: Python, R, C
     - Frameworks: TensorFlow, PyTorch, Keras, Scikit-learn, Streamlit
     - Gen AI: Langchain, Llama Index, LLMs, LLMOps
     - AI/ML: Machine Learning, Deep Learning, Natural Language Processing, Generative AI
     - Data Science: Data Analysis, Data Visualization, Predictive Modeling
     - Cloud Platforms: Cosmo Cloud, AWS, Google Cloud
     - Tools: Git, VS Code, Jupyter Notebooks, Groq
   - If neither condition is met, return "Fail."
     output "Pass" or "Fail." Do include additional explanations like missing skills and recommendation.
    """
    response = model.generate_content([prompt])
    
    return response.text

file_path_companies = 'Copy of in_data(1) (1).xlsx' 
companies_data = pd.read_excel(file_path_companies)

print("Available columns in the dataset:")
print(companies_data.columns)
user_skills_input = input("Enter your skills (comma-separated): ")
user_skills = [skill.strip() for skill in user_skills_input.split(",")]

recommended_companies = recommend_companies(companies_data, user_skills, skill_match_threshold=0.5)
if not recommended_companies.empty:
    print("\nHere are your recommended job opportunities:\n")
    detailed_recommendations = format_detailed_recommendations(recommended_companies, user_skills)
    print(detailed_recommendations)
    
    try:
        apply_choice = int(input(f"Enter the number of the job you want to apply for (1 to {len(recommended_companies)}): "))
        
        if 1 <= apply_choice <= len(recommended_companies):
            selected_company = recommended_companies.iloc[apply_choice - 1]
            print(f"You have selected: {selected_company['Profile']} at {selected_company['Company']}")
            
            company_skills = selected_company['Skills Required']
            
            resume_file = input("Enter the path to your resume file (e.g., 'resume.docx'): ")
            
            resume_file_content = extract_text_from_docx(resume_file)

            ats_comparison_report = compare_resume_to_job(resume_file_content, company_skills)

            print("\nATS Comparison Report:\n")
            print(ats_comparison_report)
        else:
            print(f"Invalid choice. Please select a valid job number between 1 and {len(recommended_companies)}.")
    except ValueError:
        print("Invalid input. Please enter a valid number.")
else:
    print("No companies match your skills.")

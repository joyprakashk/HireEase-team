import pandas as pd

# Load the dataset
file_path = r'STD_data (1).xlsx'  # Adjust the file path if necessary
data = pd.read_excel(file_path)

# Standardize column names
data.columns = data.columns.str.strip().str.lower()

def filter_candidates(data, institution=None, skills=None, area_of_interest=None, skill_match_threshold=0.5):
    """
    Filters candidates based on the given criteria.
    
    Parameters:
        data (DataFrame): The dataset containing candidate information.
        institution (str): The institution name to filter by.
        skills (list of str): A list of skills to match.
        area_of_interest (str): Area of interest to filter by.
        skill_match_threshold (float): Minimum percentage of skills that must match (0 to 1).
    
    Returns:
        DataFrame: Filtered candidates.
    """
    filtered_data = data.copy()

    # Filter by institution if provided
    if institution:
        filtered_data = filtered_data[
            filtered_data['institution'].str.contains(institution, case=False, na=False)
        ]
    
    # Filter by area of interest if provided
    if area_of_interest:
        filtered_data = filtered_data[
            filtered_data['area of interest'].str.contains(area_of_interest, case=False, na=False)
        ]

    # Filter by skills if provided
    if skills:
        skills = [skill.strip().lower() for skill in skills]
        total_skills = len(skills)
        if total_skills > 0:
            filtered_data = filtered_data[
                filtered_data['skills'].apply(
                    lambda x: (
                        sum(skill in x.lower() for skill in skills) / total_skills
                    ) >= skill_match_threshold if pd.notnull(x) else False
                )
            ]

    return filtered_data

# Example: Prompting user for inputs
institution_input = input("Enter institution name (or press Enter to skip): ").strip()
skills_input = input("Enter skills (comma-separated, or press Enter to skip): ").strip().split(',') if input("Do you want to filter by skills? (yes/no): ").lower() == "yes" else None
area_of_interest_input = input("Enter area of interest (or press Enter to skip): ").strip()

# Call the filtering function with a 50% skill match threshold
filtered_candidates = filter_candidates(
    data,
    institution=institution_input if institution_input else None,
    skills=skills_input if skills_input else None,
    area_of_interest=area_of_interest_input if area_of_interest_input else None,
    skill_match_threshold=0.5  # Match at least 50% of skills
)

# Display results
if not filtered_candidates.empty:
    print("\nFiltered Candidates:")
    print(filtered_candidates[['name', 'institution', 'skills', 'area of interest']])
else:
    print("\nNo candidates found matching the criteria.")

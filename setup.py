from setuptools import setup, find_packages

setup(
    name="Python_Python_projects", 
    version="0.1.0", 
    author="Your Name",
    author_email="your.email@example.com",
    description="A small example package",
    long_description="This is a simple example Python package that demonstrates basic setup functionality.",
    long_description_content_type="text/plain",
    packages=find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',  
    install_requires=[
        'scipy>=1.4', 
        'flask>=2.0'
    ],
)

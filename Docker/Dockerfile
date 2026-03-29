# Use official Python image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Install Jupyter (includes IPython)
RUN pip install --no-cache-dir notebook

# Expose notebook port
EXPOSE 8888

# Run Jupyter Notebook
CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root", "--NotebookApp.token=''"]

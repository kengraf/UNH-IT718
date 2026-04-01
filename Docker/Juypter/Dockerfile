# Use the desired Python slim image as the base
FROM python:slim

# Set the working directory in the container
WORKDIR /app

# Copy the rest of your application code
COPY . .

# Install the Python packages
# Using --no-cache-dir reduces the image size
RUN pip install --no-cache-dir -r requirements.txt

# Expose notebook port
EXPOSE 8888

# Run Jupyter Notebook
CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root"]

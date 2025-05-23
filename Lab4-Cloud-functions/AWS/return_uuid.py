import json
import uuid

def lambda_handler(event, context):
    path = event.get('rawPath', '/')
    if path == '/favicon.ico':
        return {
            "statusCode": 404,
            "body": "Not Found"
        }
    return(get_response())

def get_response():
    # Generate a UUID
    generated_uuid = str(uuid.uuid4())
    
    # Create HTML response
    html_content = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>Welcome</title>
    </head>
    <body>
        <h1>Hello, World!</h1>
        <p>Your session UUID: {generated_uuid}</p>
    </body>
    </html>
    """
  
    # Return the response with the Set-Cookie header
    return {
        "statusCode": 200,
        "headers": {
            "Content-Type": "text/html",
            "Set-Cookie": f"test-id={generated_uuid}; Path=/; HttpOnly"
        },
        "body": html_content
    }

if __name__ == '__main__':
     print(lambda_handler(0,0))

# AI Application in Golang

This project is an AI application developed using Golang, leveraging models like Llama 3.1, Llama 2, and Ollama for language processing and understanding. The application is designed for testing and studying AI capabilities in Golang, particularly focusing on document-based training and user interaction with models through HTTP routes.

## Project Overview

The application has four main routes:

1. **Health Check (`GET /health`)**: This route checks if the application is running correctly. It simply returns a status message to indicate that the application is live.

2. **Prompt Interaction (`POST /prompt`)**: This route accepts a JSON payload with a prompt, interacting directly with the AI model to provide a response. Example usage:
    ```bash
    curl -X POST http://localhost:8080/prompt -H "Content-Type: application/json" -d '{"prompt": "What is the capital of France?"}'
    ```

3. **Document Feed (`POST /feed`)**: This route allows uploading `.txt` files to be used as training data for the model. The content is appended to a central file (`aggregate.txt`) which acts as an aggregate document for training purposes. The example code ensures only `.txt` files are supported for security and consistency. It would be possible to further optimize this using a vector database in the future.

4. **Model Training (`POST /trainModel`)**: This route adds new training data to the aggregate file in a Q&A format. Example usage:
    ```bash
    curl -X POST http://localhost:8080/trainModel -H "Content-Type: application/json" -d '{"question": "What is the color of the sky?", "answer": "The sky is blue."}'
    ```
   This route requires a question and an answer as input, appending them to the `aggregate.txt` file.

## Why Use Golang Over Python?

While Python is the most popular language for AI and machine learning, Golang offers several advantages, especially for building AI-driven web services:

- **Performance**: Golang is a compiled language, offering superior performance compared to Python, which is interpreted. This makes it suitable for applications that require low latency and high throughput.
- **Concurrency**: Golang’s built-in concurrency model using goroutines and channels allows efficient handling of multiple requests, making it a great choice for web servers and APIs, which are essential components of AI systems.
- **Simplicity and Efficiency**: Golang has a simple syntax and is efficient to deploy, making it an ideal choice for developers who want to build scalable AI services without dealing with the complexities of more extensive ecosystems like Python’s.

## AI Models Used

### 1. Llama 3.1 (Hugging Face)

Llama 3.1 is a powerful language model that can be fine-tuned and trained for specific tasks. It is available on [Hugging Face](https://huggingface.co/) and provides state-of-the-art performance for natural language understanding tasks.

- **Performance**: It offers high accuracy and fluency but requires significant computational resources for fine-tuning and inference.
- **Isolation**: Llama 3.1, when deployed locally, can operate in isolation, meaning no internet connection is required to process requests, aligning with the application’s goal to maintain a secure and self-contained environment.

### 2. Llama 2 (Hugging Face)

Llama 2 is another language model available on Hugging Face, which serves as an earlier version of Llama 3.1. It is also highly effective but may not match the latest performance and optimization improvements of Llama 3.1.

- **Performance**: Slightly lower than Llama 3.1, but still effective for most common NLP tasks.
- **Compatibility**: It provides a good balance between resource usage and performance, making it a viable option when computational resources are limited.

### 3. Ollama

Ollama is an online platform offering language models similar to those on Hugging Face. The key difference lies in its cloud-based services, which can be beneficial for rapid prototyping and testing but may pose security and privacy concerns for production applications.

- **Performance**: The performance can vary based on network latency and internet speed. While it offers convenience, it might not be as fast or reliable as locally hosted models like Llama 3.1.
- **Internet Dependency**: Ollama requires an internet connection, which contrasts with the application’s aim to function in an isolated environment. As such, it is primarily used for experimentation and comparison.

## Future Considerations

- **Integration with Vector Databases**: The code currently aggregates documents into a single file. Future development could involve integrating a vector database to efficiently store and query documents, enhancing the model’s capability to respond based on context.
- **Optimization for Multi-model Deployment**: The application architecture is flexible and can be expanded to accommodate different models simultaneously, offering comparative analysis and improved user experience.

## Conclusion

This AI application in Golang serves as a robust and high-performance platform for AI experimentation and development. By isolating the models from the internet, it ensures security and reliability, while Golang’s efficiency and concurrency make it a powerful alternative to traditional Python-based AI systems.


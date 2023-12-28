# go-openai
A wrapper around the openai api for golang

### TODOS:
- [X] Implement GPT-4 ASAP
    - [ ] GPT-4 Vision.
    - [ ] Text Completions
    - [ ] Assistant Models
    - [ ] Function calling
    - [ ] Text to speach
    - [ ] Speech to text (Would be so dope to talk to it and have it respond
    either calling a function to perform a task etc..)

- [ ] Might want to add some sort of persistence layer
- [ ] Setup settings parsing and processing.
- [ ] Add support for image generation
- [ ] Add support for transcription generation
- [ ] Add support for translation generation
- [ ] Add support for uploading fine tuning documents.
- [ ] Input validation, ignoring misc input ie scroll characters etc..
- [ ] SSE Stream support for chat completion.
- [ ] Format the response type to make it easier to copy responses.
- [ ] Optimize the token settings so that i'm getting bigger/better responses.


## Inspiration:
- [Calling Functions with Chat models](https://cookbook.openai.com/examples/how_to_call_functions_with_chat_models)
- [What are Embeddings?](https://platform.openai.com/docs/guides/embeddings/what-are-embeddings)
- [Vector Databases and Embeddings](https://platform.openai.com/docs/guides/embeddings/how-can-i-retrieve-k-nearest-embedding-vectors-quickly)
- [OpenAI Cookbook Vector Databases](https://cookbook.openai.com/examples/vector_databases/readme)
- [Token Counting With Tiktoken (Python)](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)



### Scratch notes
// Headers
// Authorization: Bearer
// OpenAI-Organization: org-xxxx (if the user wants to attribute these requests to the organization)

// Possible routes.
* 

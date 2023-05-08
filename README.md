# Twitter Handle Checker

This is a simple Go program that checks the availability of a specified Twitter handle every hour. When the handle becomes available, it sends a Pushover notification via HTTP.

## How to use

1. Clone this repository:

```bash
git clone https://github.com/janyksteenbeek/twitter-handle-checker.git
````

2. Change to the `twitter-handle-checker` directory:
```bash
cd twitter-handle-checker
```

3. Build the Docker image:
```bash
docker build -t twitter-handle-checker .
```

4. Run the Docker container, passing your desired Twitter handle and Pushover token as environment variables:

```bash
docker run -d --name twitter-handle-checker -p 8080:8080 \
  -e CHECKER_USERNAME=desired_username \
  -e CHECKER_PUSHOVER_TOKEN=your_pushover_token \
  twitter-handle-checker
```
Replace desired_username and your_pushover_token with the actual values you want to use.

The Go program inside the container will check the availability of the Twitter handle every hour and send a Pushover notification when it becomes available.

# Security
If you discover any security-related issues, please email me@janyk.dev instead of using the issue tracker. All security vulnerabilities will be promptly addressed.

# License
This project is open-sourced software licensed under the [MIT license](LICENSE.md). Please see the [license file](LICENSE.md) for more information.
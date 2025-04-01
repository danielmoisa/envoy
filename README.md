# Envoy

A modern job application tracking system.

## Table of Contents

1.  [Installation](#installation)
2.  [Usage](#usage)
3.  [Contributing](#contributing)
4.  [License](#license)

## Installation

To install Envoy, follow these steps:

1.  Clone the repository:

    ```bash
    git clone https://github.com/danielmoisa/envoy.git
    ```

2.  Navigate to the project directory:

    ```bash
    cd envoy
    ```

3.  Install the dependencies:

    ```bash
    go mod tidy
    ```

4.  Configure the environment variables:

    *   Create a `.env` file in the root directory.
    *   Add the necessary environment variables (e.g., database connection details, API keys).

5.  Create database:

    ```bash
    docker-compose up -d
    ```

6.  Run the application:

    ```bash
    make start
    ```


The application will start on the configured port (default is 8080).

## Contributing

We welcome contributions to Envoy! To contribute, follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes.
4.  Test your changes thoroughly.
5.  Submit a pull request.

## License

Envoy is licensed under the [MIT License](LICENSE).

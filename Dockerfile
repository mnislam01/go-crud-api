FROM golang:1.18 as GolangEnvironement
# Set a working directory
WORKDIR /app
# Install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
copy . .
# Export the PORT
EXPOSE 8000
# Compile the codes
RUN go build -o app .
# Run the binary
CMD ./app

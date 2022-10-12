FROM golang:1.16-alpine

WORKDIR /app

# Copy source code  
RUN mkdir outputs
RUN mkdir pkg
COPY pkg pkg 
COPY outputs outputs
COPY go.mod go.sum main.go ./

# Build the code 
RUN go build

# Expose server port 
EXPOSE 8080

# Run the executable
CMD ["./url-shortner"]


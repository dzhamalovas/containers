# ---------- GO BUILD ----------
    FROM golang:1.21-alpine AS builder

    WORKDIR /app
    RUN apk add --no-cache gcc musl-dev sqlite-dev
    
    COPY app/go.mod ./
    RUN go mod download
    
    COPY app/ .
    RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o credit-risk-app .    
    # ---------- RUNTIME ----------
    FROM python:3.11-alpine
    
    WORKDIR /app
    
    RUN apk add --no-cache sqlite
    
    COPY --from=builder /app/credit-risk-app .
    COPY streamlit_app/ ./streamlit_app/
    COPY .env .
    RUN pip install --no-cache-dir -r streamlit_app/requirements.txt
    
    RUN mkdir data
    
    EXPOSE 8501
    
    CMD ["sh", "-c", "./credit-risk-app && streamlit run streamlit_app/app.py --server.port=8501 --server.address=0.0.0.0"]
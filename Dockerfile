FROM swaggerapi/swagger-codegen-cli-v3

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

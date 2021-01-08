TODO:
- [ ] fix https://github.com/serverless/serverless/pull/7619 
- [ ] template modifications
- [ ] tests
# SDK Generator

This GitHub Action syncs an OpenAPI spec with client SDK repositories.

## Inputs

### `access-token`

**Required** Provide a GitHub App access token with create and write permissions.

## Outputs

### `exit-code`

The program exit code.

## Example usage

```
uses: actions/SDK-Generator@v1
with:
  env:
    - access-token: {{ github.secrets.SDK_GENERATOR_ACCESS_TOKEN }}
```

Set `SDK_GENERATOR_ACCESS_TOKEN` in your repository secrets

## Optional: AWS API Gateway Integration

If you want to have the SDK updated by using API Gateway:

1. Create an IAM policy
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ExportSwagger",
            "Effect": "Allow",
            "Action": "apigateway:GET",
            "Resource": "arn:aws:apigateway:*::/restapis/{apiId}/stages/dev/exports/oas30"
        }
    ]
}
```
2. Attach policy to user or role
3. Supply access key & access token as secrets
```yaml
...
with:
  env:
    ...
    - aws-key-id: {{ github.secrets.AWS_ACCESS_KEY_ID }}
    - aws-key-secret: {{ github.secrets.AWS_SECRET_ACCESS_KEY }}
...
```

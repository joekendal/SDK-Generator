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
```

Set SDK\_GENERATOR\_ACCESS\_TOKEN in your repository secrets

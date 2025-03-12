# Nopeman

Nopeman "suspects" your Postman collection contains sensitive data.  
Nopeman checks the collection before you share it with confidence.

![Nopeman in action](screen.gif)

## Requirements

- [x] First install [jp](https://github.com/jmespath/jp#installing)

## Usage

```bash
nopeman <inputfile>
```

## Sample output

![Sample output](sample_run.png)

## Configuration

Edit the `redact.yml` file to add/remove rules.

```yaml
- name: Display name of the rule
  match: "[]" # JMESPath query to match the data you want to redact
```

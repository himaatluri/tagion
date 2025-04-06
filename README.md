# Tagion

A command-line tool for managing tags in AWS CloudFormation templates.

## Overview

Tagion simplifies the process of adding and managing tags across multiple AWS CloudFormation templates. It analyzes templates and automatically adds specified tags to supported AWS resources while preserving existing tags.

## Features

- Process single templates or entire directories
- Support for both YAML and JSON CloudFormation templates
- Preview changes before applying
- Smart detection of resources that support tags
- Preserve existing tags while adding new ones
- Non-destructive updates

## Installation

```bash
go install github.com/himaatluri/tagion@latest
```

## Usage

1. Create a JSON file with your tags:

```json
{
  "tags": {
    "Environment": "Production",
    "Project": "TagionCFN",
    "Owner": "DevOps",
    "ManagedBy": "Tagion"
  }
}
```

2. Run Tagion:

```bash
# For a single template
tagion -tags tags.json -path template.yaml

# For a directory of templates
tagion -tags tags.json -path templates/
```

## Supported Resources

Currently supports tagging for:
- AWS::EC2::* resources
- AWS::S3::* resources
- AWS::RDS::* resources
- AWS::DynamoDB::* resources
- AWS::Lambda::* resources

## Example Output

```
╭────────────────────────────────────────────┬───────────┬────────────────────╮
│ TEMPLATE                                   │ RESOURCES │ STATUS             │
├────────────────────────────────────────────┼───────────┼────────────────────┤
│ templates/ec2-no-tags.yaml                 │         1 │ Will be modified   │
│ templates/s3-with-tags.yaml                │         1 │ Has tags           │
│ templates/multi-resources.json             │         2 │ Will be modified   │
│ templates/unsupported.yaml                 │         2 │ No changes needed  │
├────────────────────────────────────────────┼───────────┼────────────────────┤
│ Total                                      │         6 │ To modify: 2       │
╰────────────────────────────────────────────┴───────────┴────────────────────╯
```

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Documentation

For more detailed information, check out:
- [Blog Post](docs/blog-post.md)
- [Example Templates](examples/templates/)

## Future Enhancements

- Support for more AWS resource types
- Custom tag validation rules
- Integration with AWS Organizations tag policies
- Tag removal and modification features
- CI/CD pipeline integration
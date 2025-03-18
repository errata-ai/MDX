# MDX

A Vale configuration for sites using [MDX][1]. You'll want to following the
[MDX installation][2] instructions before using this package.

## Getting Started

To get started, add the package to your configuration file (as shown below) and then run `vale sync`.

```ini
StylesPath = styles
MinAlertLevel = suggestion

Packages = MDX

[*]
# The 'Vale' is optional;
#
# Run `ls-config` to see the
# MDX configuration settings.
BasedOnStyles = Vale
```

## Testing

```
$ make test
```

[1]: https://mdxjs.com
[2]: https://vale.sh/docs/formats/mdx

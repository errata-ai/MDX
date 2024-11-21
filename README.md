# MDX

A Vale configuration for sites using [MDX][1]:

- [x] Adds support for using MDX comment syntax (`{/* ... */}`).
- [x] Ignores `import`s and `export`s.
- [x] Ignores `<Component ... />` and `<Component> ... </Component>`.
- [x] Ignores `{ ... }`.

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

# TOML and ENV files
Toml files and env files are both related to saving configuration files. But they vary in some aspects:

**TOML:**

- Used for complex configuration settings.
- Has a structured and hierarchical format with nested data.
- Human-readable and easy to understand.
- Suitable for applications or systems with multiple configuration sections.
  
**ENV (Environment) files:**

- Used for simple configuration settings, especially environment variables.
- Has a flat format with individual key-value pairs.
- Human-readable but it lacks hierarchical structure.
- Suitable for managing environment-specific settings without modifying code.

 If we need to manage simple key-value pairs, especially for environment-specific settings, ENV are always preferred. On the other hand, if we require a more structured and hierarchical configuration setup, TOML files provide a better option. 

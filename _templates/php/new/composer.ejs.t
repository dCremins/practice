---
to: <%= h.changeCase.pascal(name) %>/composer.json
---
{
    "name": "php/<%= h.changeCase.snakeCase(name) %>",
    "description": "<%= h.changeCase.pascal(name) %>",
    "authors": [
        {
            "name": "Devin Cremins",
            "email": "devincrem@gmail.com"
        }
    ],
    "autoload": {
        "psr-4": {
            "Practice\\": "."
        }
    },
    "require-dev": {
        "phpunit/phpunit": "^9.5"
    }
}
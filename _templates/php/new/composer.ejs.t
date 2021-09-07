---
to: <%= h.changeCase.pascal(name) %>/composer.json
---
{
    "name": "parctice/php/<%= h.changeCase.snakeCase(name) %>",
    "description": "<%= h.changeCase.pascal(name) %>",
    "authors": [
        {
            "name": "Devin Cremins",
            "email": "devincrem@gmail.com"
        }
    ],
    "require-dev": {
        "phpunit/phpunit": "^7.1"
    }
}
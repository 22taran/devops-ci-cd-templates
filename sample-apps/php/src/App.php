<?php

namespace App;

class App
{
    public function greet(?string $name = null): string
    {
        if ($name === null || trim($name) === '') {
            return 'Hello, World!';
        }
        return "Hello, {$name}!";
    }

    public function add(int $a, int $b): int
    {
        return $a + $b;
    }
}

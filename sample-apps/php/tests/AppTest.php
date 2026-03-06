<?php

use PHPUnit\Framework\TestCase;
use App\App;

class AppTest extends TestCase
{
    public function testGreetWithName(): void
    {
        $app = new App();
        $this->assertEquals('Hello, DevOps!', $app->greet('DevOps'));
    }

    public function testGreetWithNull(): void
    {
        $app = new App();
        $this->assertEquals('Hello, World!', $app->greet(null));
    }

    public function testGreetWithEmpty(): void
    {
        $app = new App();
        $this->assertEquals('Hello, World!', $app->greet(''));
    }

    public function testAdd(): void
    {
        $app = new App();
        $this->assertEquals(5, $app->add(2, 3));
    }
}

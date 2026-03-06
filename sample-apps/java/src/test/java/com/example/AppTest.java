package com.example;

import org.junit.Test;
import static org.junit.Assert.*;

public class AppTest {

    @Test
    public void testGreetWithName() {
        App app = new App();
        assertEquals("Hello, DevOps!", app.greet("DevOps"));
    }

    @Test
    public void testGreetWithNull() {
        App app = new App();
        assertEquals("Hello, World!", app.greet(null));
    }

    @Test
    public void testGreetWithEmpty() {
        App app = new App();
        assertEquals("Hello, World!", app.greet(""));
    }

    @Test
    public void testAdd() {
        App app = new App();
        assertEquals(5, app.add(2, 3));
    }
}

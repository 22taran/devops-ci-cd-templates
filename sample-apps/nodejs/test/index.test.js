const { add, greet } = require('../src/index');

describe('add', () => {
    test('adds two numbers', () => {
        expect(add(2, 3)).toBe(5);
    });

    test('adds negative numbers', () => {
        expect(add(-1, 1)).toBe(0);
    });
});

describe('greet', () => {
    test('greets with name', () => {
        expect(greet('DevOps')).toBe('Hello, DevOps!');
    });

    test('greets with default when empty', () => {
        expect(greet('')).toBe('Hello, World!');
    });

    test('greets with default when null', () => {
        expect(greet(null)).toBe('Hello, World!');
    });
});

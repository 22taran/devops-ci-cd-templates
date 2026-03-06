const express = require('express');
const app = express();
const PORT = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.json({ message: 'Hello, World!' });
});

app.get('/health', (req, res) => {
  res.json({ status: 'healthy', timestamp: new Date().toISOString() });
});

function add(a, b) {
  return a + b;
}

function greet(name) {
  if (!name || name.trim() === '') {
    return 'Hello, World!';
  }
  return `Hello, ${name}!`;
}

if (require.main === module) {
  app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
  });
}

module.exports = { app, add, greet };

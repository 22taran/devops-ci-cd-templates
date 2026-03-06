require_relative '../app'

RSpec.describe 'App' do
  describe '#greet' do
    it 'greets with a name' do
      expect(greet('DevOps')).to eq('Hello, DevOps!')
    end

    it 'returns default when nil' do
      expect(greet(nil)).to eq('Hello, World!')
    end

    it 'returns default when empty' do
      expect(greet('')).to eq('Hello, World!')
    end
  end

  describe '#add' do
    it 'adds two numbers' do
      expect(add(2, 3)).to eq(5)
    end
  end
end

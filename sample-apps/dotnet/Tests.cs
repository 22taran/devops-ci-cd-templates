using Xunit;

namespace SampleApp.Tests;

public class AppTests
{
    [Fact]
    public void Greet_WithName_ReturnsGreeting()
    {
        var app = new App();
        Assert.Equal("Hello, DevOps!", app.Greet("DevOps"));
    }

    [Fact]
    public void Greet_WithNull_ReturnsDefault()
    {
        var app = new App();
        Assert.Equal("Hello, World!", app.Greet(null));
    }

    [Fact]
    public void Greet_WithEmpty_ReturnsDefault()
    {
        var app = new App();
        Assert.Equal("Hello, World!", app.Greet(""));
    }

    [Fact]
    public void Add_TwoNumbers_ReturnsSum()
    {
        var app = new App();
        Assert.Equal(5, app.Add(2, 3));
    }
}

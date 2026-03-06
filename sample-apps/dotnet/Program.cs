namespace SampleApp;

public class App
{
    public string Greet(string? name)
    {
        if (string.IsNullOrWhiteSpace(name))
            return "Hello, World!";
        return $"Hello, {name}!";
    }

    public int Add(int a, int b) => a + b;

    public static void Main(string[] args)
    {
        var app = new App();
        Console.WriteLine(app.Greet("DevOps"));
    }
}

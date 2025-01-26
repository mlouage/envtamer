namespace envtamer.Data;

using Microsoft.EntityFrameworkCore;

public class EnvTamerContext : DbContext
{
    public DbSet<EnvVariable> EnvVariables { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
        var dbPath = Path.Combine(
            Environment.GetFolderPath(Environment.SpecialFolder.UserProfile),
            ".envtamer",
            "envtamer.db"
        );
        optionsBuilder.UseSqlite($"Data Source={dbPath};Mode=ReadWriteCreate");
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<EnvVariable>()
            .HasKey(e => new { e.Directory, e.Key });
    }
}

public class EnvVariable
{
    public required string Directory { get; set; }
    public required string Key { get; set; }
    public required string Value { get; set; }
}

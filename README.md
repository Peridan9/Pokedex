
# 🔍 Pokedex CLI 📜

This project is a **command-line-based Pokedex** that interacts with the **PokeAPI** to fetch and display Pokémon-related information.  
It allows users to explore Pokémon locations, catch Pokémon with a probability-based system, and view their captured Pokémon list.  

Developed in **Go**, this project includes features such as **caching API responses**, **pagination**, and a **probabilistic catching system** based on a Pokémon’s base experience.  

---

## ✨ Features

- **Explore Locations:** Fetch and list Pokémon found in a given location.
- **Catch Pokémon:** Attempt to catch Pokémon, with success probability based on their base experience.
- **View Caught Pokémon:** Keep track of all Pokémon you’ve caught using the `pokedex` command.
- **Location Pagination:** Browse Pokémon locations with `map` (next 20) and `mapb` (previous 20).
- **Caching System:** API responses are cached to improve performance and reduce redundant API calls.
- **Command-Line Interface (CLI):** Simple text-based input system for interaction.

---

## 🛠️ Technologies Used

- **Go**: Main language for CLI development.
- **PokeAPI**: External REST API for fetching Pokémon data.
- **Go HTTP Client**: Used for API calls.
- **Caching Mechanism**: Implements an in-memory cache with automatic cleanup.

---

## 📦 Installation & Setup

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/Peridan9/Pokedex.git
   cd Pokedex
   ```

2. **Build the Project**  
   ```bash
   go build -o pokedex
   ```

3. **Run the Pokedex CLI**  
   ```bash
   ./pokedex
   ```
---

## 🎮 Available Commands

| Command              | Description |
|----------------------|-------------|
| `help`              | Displays a help message with available commands. |
| `map`               | Fetches the next 20 locations. |
| `mapb`              | Fetches the previous 20 locations. |
| `explore <location>` | Lists all Pokémon encounters in a given location. |
| `catch <pokemon>`   | Attempts to catch a Pokémon (success probability based on base experience). |
| `pokedex`           | Displays a list of caught Pokémon. |
| `exit`              | Exits the Pokedex CLI. |

---

## 📂 File Structure

```
Pokedex/
├── internal/
│   ├── pokeapi/            # API client and data fetching logic
│   │   ├── client.go
│   │   ├── location_list.go
│   │   ├── pokeapi.go
│   │   ├── types_locations.go
│   │   ├── types_pokemon.go
│   ├── pokecache/          # Caching system
│       ├── cache.go
│── main.go                 # Entry point for the CLI
│── command_explore.go       # Logic for 'explore' command
│── command_catch.go         # Logic for 'catch' command
│── command_pokedex.go       # Logic for 'pokedex' command
│── README.md                # This file
│── go.mod                   # Go module dependencies
```

---

## ⚔️ Catching Pokémon Logic

- Each Pokémon has a **base experience** value.
- The **higher the base experience**, the **harder it is to catch**.
- The catch probability is calculated using:
  ```go
  res := rand.Intn(pokemon.BaseExperience)
  if res < 40 {
      fmt.Println("You caught the Pokémon!")
  } else {
      fmt.Println("The Pokémon escaped!")
  }
  ```
- Pokémon are stored in a **map inside `config`**, so the caught list persists **during the session**.

---

## 🚀 Example Usage

```bash
Pokedex > map
# Displays 20 Pokémon locations

Pokedex > explore mt-coronet-1f-route-207
# Lists Pokémon found in that location

Pokedex > catch pikachu
# Attempts to catch Pikachu with probability based on base experience

Pokedex > pokedex
# Shows all caught Pokémon
```

---

## 📝 Future Improvements

- Persistent storage for caught Pokémon across sessions.
- Add moves and abilities to `explore` details.
- More advanced catching mechanics (e.g., Poké Balls with different catch rates).

---

## 📡 API Reference

- **PokeAPI**: [https://pokeapi.co/](https://pokeapi.co/)
- **Location Endpoint**: `/location-area/{name}`
- **Pokémon Endpoint**: `/pokemon/{name}`

---

Let me know if you’d like any changes! 🚀🔥

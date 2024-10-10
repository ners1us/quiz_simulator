# quiz_simulator
Quiz simulator application

## Building and running the application
```bash
docker compose build && docker compose run --rm app ./quiz_simulator
```

## Notes
- All problems are listed in the ```quiz.csv``` file, you can edit it and add your own problems in the ```problem,answer``` format.
- The program offers you to set your own time for the timer
- Problems are always presented in random order in the application.
- After the set time has elapsed, the program will terminate its work.
- As a result, the program displays the number of correctly solved problems from the amount of problems presented in the ```quiz.csv``` file.

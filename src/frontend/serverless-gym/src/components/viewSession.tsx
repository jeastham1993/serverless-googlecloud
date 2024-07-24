import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  Button,
  Card,
  CardContent,
  IconButton,
  Snackbar,
  TextField,
  Typography,
} from "@mui/material";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Unstable_Grid2";
import DeleteIcon from "@mui/icons-material/Delete";
import { Session, SessionExercise } from "../models/session";
import { ExerciseHistroy } from "../models/exerciseHistory";
import ExerciseTitle from "./exerciseTitle";
import { ApiService } from "../services/apiService";
import ExerciseTile from "./exerciseTile";

const style = {
  my: 2,
};

const modalStyle = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: "75vw",
  bgcolor: "background.paper",
  boxShadow: 24,
  p: 4,
};

export default function ViewSession(props: any) {
  const { id } = useParams();
  const navigate = useNavigate();
  const [isLoading, setLoading] = useState(true);
  const [open, setOpen] = useState(false);
  const [session, setSession] = useState<Session>({
    id: "",
    date: "",
    exercises: [],
  });
  const [newExercise, setNewExercise] = useState<SessionExercise>({
    name: "",
    reps: 0,
    weight: 0,
    set: 0,
  });
  const apiService = new ApiService(navigate);

  const handleClose = (
    event: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }

    setOpen(false);
  };

  const saveSession = async () => {
    await apiService.saveSession(session);

    setOpen(true);
  };

  const finishSession = async () => {
    await apiService.finishSession(session);
  };

  const handleRepsChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>,
    exercise: SessionExercise
  ) => {
    const parsedInt = parseInt(e.target.value);

    if (isNaN(parsedInt)) {
      e.target.value = "0";
      return;
    }

    const newSession: SessionExercise = {
      name: exercise.name,
      set: exercise.set,
      weight: exercise.weight,
      reps: parseInt(e.target.value),
    };

    const newExercises = session.exercises.map((existing) => {
      if (existing.name === exercise.name && existing.set === exercise.set) {
        return newSession;
      } else {
        return existing;
      }
    });

    setSession({ ...session, exercises: newExercises });
  };

  const handleWeightChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>,
    exercise: SessionExercise
  ) => {
    var checkValue: string = e.target.value;

    if (checkValue.endsWith(".")) {
      checkValue = checkValue + "5";
    }

    const parsedInt = parseFloat(checkValue);

    if (isNaN(parsedInt)) {
      console.log("Nope");
      return;
    }

    console.log(parsedInt);

    const newSession: SessionExercise = {
      name: exercise.name,
      set: exercise.set,
      weight: parseFloat(checkValue),
      reps: exercise.reps,
    };

    const newExercises = session.exercises.map((existing) => {
      if (existing.name === exercise.name && existing.set === exercise.set) {
        return newSession;
      } else {
        return existing;
      }
    });

    setSession({ ...session, exercises: newExercises });
  };

  const handleNewExerciseRepsChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
  ) => {
    setNewExercise({ ...newExercise, reps: parseInt(e.target.value) });
  };

  const handleNewExerciseWeightChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
  ) => {
    setNewExercise({ ...newExercise, weight: parseFloat(e.target.value) });
  };

  const handleNewExerciseNameChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>
  ) => {
    setNewExercise({ ...newExercise, name: e.target.value });
  };

  const addNewExerciseToSession = (e: any) => {
    const existingExercises = session.exercises;
    const matchedExercises = existingExercises.filter(
      (e) => e.name === newExercise.name
    );

    const exercise = newExercise;

    if (matchedExercises.length > 0) {
      const exercise = newExercise;
      exercise.set = matchedExercises.length + 1;
    } else {
      exercise.set = 1;
    }

    existingExercises.push(exercise);
    existingExercises.sort((a, b) => {
      const aSortValue = `${a.name}-${a.set}`;
      const bSortValue = `${b.name}-${b.set}`;

      return aSortValue.localeCompare(bSortValue);
    });

    console.log(existingExercises);

    setSession({ ...session, exercises: existingExercises });

    saveSession();
  };

  const removeExercise = (exercise: SessionExercise) => {
    const newExercises = session.exercises.filter((existing) => {
      if (existing.name === exercise.name && existing.set === exercise.set) {
        return false;
      } else {
        return true;
      }
    });

    setSession({ ...session, exercises: newExercises });

    saveSession();
  };

  const refreshData = async () => {
    const data = await apiService.getSession(id);

    setSession(data);
    setLoading(false);
  };

  const handleOnBlur = async (evt: any) => {
    await saveSession();
  };

  useEffect(() => {
    refreshData();
  }, []);

  return (
    <div>
      <Box sx={style}>
        <Typography
          id="modal-modal-title"
          variant="h5"
          component="h5"
          sx={{ my: 2 }}
        >
          {session.id}
        </Typography>
        <Grid container>
          {session.exercises.map((e) => (
            <Grid key={`${e.name}-${e.set}`} xs={12} lg={3}>
              <ExerciseTile
                exercise={e}
                readOnly={false}
                onDeleteExercise={(e: SessionExercise) => removeExercise(e)}
                onRepsChange={(evt, e) => handleRepsChange(evt, e)}
                onWeightChange={(evt, e) => handleWeightChange(evt, e)}
                onWeightBlur={(evt) => handleOnBlur(evt)}
              />
            </Grid>
          ))}
        </Grid>
        <Card sx={{ margin: "1rem" }}>
          <CardContent>
            <Grid container>
              <Grid>
                <TextField
                  sx={{ mx: 1 }}
                  label="Exercise"
                  variant="outlined"
                  value={newExercise.name}
                  onChange={(evt) => handleNewExerciseNameChange(evt)}
                />
              </Grid>
              <Grid>
                <TextField
                  sx={{ mx: 1 }}
                  label="Reps"
                  type="number"
                  variant="outlined"
                  value={newExercise.reps}
                  onChange={(evt) => handleNewExerciseRepsChange(evt)}
                />
              </Grid>
              <Grid>
                <TextField
                  sx={{ mx: 1 }}
                  label="Weight"
                  variant="outlined"
                  type="number"
                  value={newExercise.weight}
                  onChange={(evt) => handleNewExerciseWeightChange(evt)}
                />
              </Grid>
              <Grid>
                <Button variant="outlined" onClick={addNewExerciseToSession}>
                  Add +
                </Button>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
        <Button sx={{ my: 2 }} variant="outlined" onClick={finishSession}>
          Finish Session
        </Button>
      </Box>
      <Snackbar
        open={open}
        autoHideDuration={3000}
        onClose={handleClose}
        message="Saved"
      />
    </div>
  );
}

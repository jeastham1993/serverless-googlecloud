"use client";

import { useState, useEffect } from "react";
import {
  Alert,
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
import axios from "axios";
import DeleteIcon from "@mui/icons-material/Delete";
import { useRouter } from "next/navigation";

const style = {
  my: 2,
};

export default function Page({ params }: { params: { id: string } }) {
  const router = useRouter();
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
    await axios.put(
      `https://gcloud-go-7tq7m2dbcq-nw.a.run.app/session/${session.id}`,
      session
    );

    setOpen(true);
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

  const refreshData = () => {
    console.log(params.id);
    fetch(`https://gcloud-go-7tq7m2dbcq-nw.a.run.app/session/${params.id}`)
      .then((res) => res.json())
      .then((data) => {
        setSession(data);
        setLoading(false);
      });
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
              <Card sx={{ margin: "1rem" }}>
                <CardContent>
                  <IconButton
                    aria-label="delete"
                    style={{ float: "right" }}
                    onClick={() => removeExercise(e)}
                  >
                    <DeleteIcon />
                  </IconButton>
                  <Typography>
                    {e.set}. {e.name}
                  </Typography>
                  <TextField
                    sx={{ my: 2 }}
                    label="Reps"
                    variant="outlined"
                    value={e.reps}
                    onChange={(evt) => handleRepsChange(evt, e)}
                  />
                  <TextField
                    sx={{ my: 2 }}
                    label="Weight"
                    variant="outlined"
                    value={e.weight}
                    onChange={(evt) => handleWeightChange(evt, e)}
                    onBlur={handleOnBlur}
                  />
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
        <Card sx={{ margin: "1rem" }}>
          <CardContent>
            <TextField
              sx={{ mx: 2 }}
              label="Exercise"
              variant="outlined"
              value={newExercise.name}
              onChange={(evt) => handleNewExerciseNameChange(evt)}
            />
            <TextField
              sx={{ mx: 2 }}
              label="Reps"
              variant="outlined"
              value={newExercise.reps}
              onChange={(evt) => handleNewExerciseRepsChange(evt)}
            />
            <TextField
              sx={{ mx: 2 }}
              label="Weight"
              variant="outlined"
              value={newExercise.weight}
              onChange={(evt) => handleNewExerciseWeightChange(evt)}
            />
            <Button variant="outlined" onClick={addNewExerciseToSession}>
              Add +
            </Button>
          </CardContent>
        </Card>
        <Button sx={{ my: 2 }} variant="outlined" onClick={saveSession}>
          Save Session
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

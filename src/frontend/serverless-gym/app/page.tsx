"use client";

import { useState, useEffect } from "react";
import {
  Button,
  Fab,
  Modal,
  Paper,
  styled,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
  Typography,
} from "@mui/material";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Unstable_Grid2"; // Grid version 2
import Divider from "@mui/material/Divider";
import AddIcon from "@mui/icons-material/Add";
import axios from "axios";
import { useRouter } from "next/navigation";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: "75vw",
  bgcolor: "background.paper",
  boxShadow: 24,
  p: 4,
};

export default function Home() {
  const router = useRouter();
  const [data, setData] = useState<Workout[]>([]);
  const [isLoading, setLoading] = useState(true);
  const [newSessionName, setNewSessionName] = useState("");
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => {
    refreshData();
    setOpen(false);
  };
  const [newWorkout, setNetWorkout] = useState<Workout>({
    id: "",
    name: "",
    exercises: [],
  });
  const [newExercise, setNewExercise] = useState<Exercise>({
    name: "",
    sets: 0,
    reps: 0,
  });

  const handleNameChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) =>
    setNetWorkout({ ...newWorkout, name: e.target!.value });

  const handleSetsChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    setNewExercise({ ...newExercise, sets: parseInt(e.target.value) });
    console.log(newExercise);
  };

  const handleRepsChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    setNewExercise({ ...newExercise, reps: parseInt(e.target.value) });
  };

  const handleExerciseNameChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    setNewExercise({ ...newExercise, name: e.target.value });
  };

  const handleNewSessionNameChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    setNewSessionName(e.target.value);
  };

  const saveExercise = () => {
    console.log(newExercise);
    newWorkout.exercises.push({
      name: newExercise.name,
      reps: newExercise.reps,
      sets: newExercise.sets,
    });

    setNetWorkout(newWorkout);

    console.log(newWorkout);

    setNewExercise({
      name: "",
      sets: 0,
      reps: 0,
    });
  };

  const startWorkoutFromSession = async (workoutId: string) => {
    if (newSessionName.length === 0) {
      return;
    }

    const postResponse = await axios.post<Session>(
      "https://gcloud-go-7tq7m2dbcq-nw.a.run.app/session/from",
      {
        workoutId,
      }
    );

    router.push(`sessions/${encodeURIComponent(postResponse.data.id)}`);
  };

  const saveWorkout = async () => {
    console.log(newWorkout);

    const postResponse = await axios.post(
      "https://gcloud-go-7tq7m2dbcq-nw.a.run.app/workout",
      newWorkout
    );

    handleClose();
  };

  const refreshData = () => {
    fetch("https://gcloud-go-7tq7m2dbcq-nw.a.run.app/workout")
      .then((res) => res.json())
      .then((data) => {
        if (data === null) {
          data = [];
        }

        setData(data);
        setLoading(false);
      });
  };

  useEffect(() => {
    refreshData();
  }, []);

  return (
    <div>
      <Box marginBlockStart={4}>
        <Grid container>
          <Grid xs={12}>
            {data.map((d) => (
              <Box key={d.id} sx={{ minWidth: 275, p: 2 }}>
                <Typography variant="h3">{d.name}</Typography>
                <TableContainer component={Paper}>
                  <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell>Name</TableCell>
                        <TableCell align="right">Sets</TableCell>
                        <TableCell align="right">Reps</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {d.exercises.map((e) => (
                        <TableRow
                          key={e.name}
                          sx={{
                            "&:last-child td, &:last-child th": { border: 0 },
                          }}
                        >
                          <TableCell component="th" scope="row">
                            {e.name}
                          </TableCell>
                          <TableCell align="right">{e.reps}</TableCell>
                          <TableCell align="right">{e.sets}</TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
                <TextField
                  sx={{ my: 2, mx: 1 }}
                  label="Session Name"
                  variant="outlined"
                  value={newSessionName}
                  onChange={handleNewSessionNameChange}
                />
                <Button
                  variant="contained"
                  sx={{ my: 2 }}
                  onClick={() => startWorkoutFromSession(d.id)}
                >
                  Start Session
                </Button>
                <Divider sx={{ my: 2 }} />
              </Box>
            ))}
          </Grid>
        </Grid>
      </Box>
      <Fab
        sx={{ position: "absolute", right: 10, bottom: 10 }}
        color="primary"
        aria-label="add"
        onClick={handleOpen}
      >
        <AddIcon />
      </Fab>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography
            id="modal-modal-title"
            variant="h5"
            component="h5"
            sx={{ my: 2 }}
          >
            Create a new workout
          </Typography>
          <TextField
            sx={{ my: 2 }}
            label="Name"
            variant="outlined"
            value={newWorkout.name}
            onChange={handleNameChange}
          />
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Sets</TableCell>
                  <TableCell>Reps</TableCell>
                  <TableCell></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {newWorkout.exercises.map((e) => (
                  <TableRow
                    key={e.name}
                    sx={{
                      "&:last-child td, &:last-child th": { border: 0 },
                    }}
                  >
                    <TableCell component="th" scope="row">
                      {e.name}
                    </TableCell>
                    <TableCell>{e.sets}</TableCell>
                    <TableCell>{e.reps}</TableCell>
                    <TableCell></TableCell>
                  </TableRow>
                ))}
                <TableRow
                  sx={{
                    "&:last-child td, &:last-child th": { border: 0 },
                  }}
                >
                  <TableCell component="th" scope="row">
                    <TextField
                      sx={{ my: 2 }}
                      label="Name"
                      variant="outlined"
                      value={newExercise.name}
                      onChange={handleExerciseNameChange}
                    />
                  </TableCell>
                  <TableCell>
                    <TextField
                      sx={{ my: 2 }}
                      label="Sets"
                      variant="outlined"
                      value={newExercise.sets}
                      onChange={handleSetsChange}
                    />
                  </TableCell>
                  <TableCell>
                    <TextField
                      sx={{ my: 2 }}
                      label="Reps"
                      variant="outlined"
                      value={newExercise.reps}
                      onChange={handleRepsChange}
                    />
                  </TableCell>
                  <TableCell>
                    <Button variant="outlined" onClick={saveExercise}>
                      Save
                    </Button>
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </TableContainer>
          <Button sx={{ my: 2 }} variant="outlined" onClick={saveWorkout}>
            Save Workout
          </Button>
        </Box>
      </Modal>
    </div>
  );
}

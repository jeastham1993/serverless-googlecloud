import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
  Button,
  Fab,
  Link,
  Modal,
  Paper,
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
import { initializeApp } from "firebase/app";
import { firebaseConfig } from "../config";
import { getAuth } from "firebase/auth";
import { isAuthenticated } from "../services/authService";
import { Session, SessionExercise } from "../models/session";
import ExerciseTitle from "./exerciseTitle";
import { ApiService } from "../services/apiService";

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

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

export default function SessionPage() {
  const navigate = useNavigate();
  const [data, setData] = useState<Session[]>([]);
  const [isLoading, setLoading] = useState(true);
  const [open, setOpen] = useState(false);
  const [newSessionName, setNewSessionName] = useState("");
  const [session, setSession] = useState<Session>({
    id: "",
    date: "",
    exercises: [],
    status: ""
  });

  const handleOpen = () => setOpen(true);
  const handleClose = () => {
    setSession({ id: "", date: "", exercises: [], status: "" });
    setOpen(false);
  };
  const apiService = new ApiService(navigate);

  const saveSession = async () => {
    await apiService.saveSession(session);

    setSession({
      id: "",
      date: "",
      exercises: [],
      status: ""
    });

    refreshData();

    handleClose();
  };

  const viewSession = (session: Session) => {
    setSession(session);

    handleOpen();
  };

  const startSessionFromSession = async (sessionId: string) => {
    if (newSessionName.length === 0) {
      return;
    }

    const postResponse = await apiService.duplicateSession(
      sessionId,
      newSessionName
    );

    navigate(`${encodeURIComponent(postResponse.id)}`);
  };

  const handleNewSessionNameChange = (e: any) => {
    setNewSessionName(e.target.value);
  };

  const handleRepsChange = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>,
    exercise: SessionExercise
  ) => {
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
    const newSession: SessionExercise = {
      name: exercise.name,
      set: exercise.set,
      weight: parseInt(e.target.value),
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

  const refreshData = async () => {
    const data = await apiService.listSessions();

    setData(data);
    setLoading(false);
  };

  const getLink = (session: Session) => {
    const link = `session/${encodeURIComponent(session.id)}`;
    return link;
  };

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate("/login");
    }
    refreshData();
  }, []);

  return (
    <div>
      <Box marginBlockStart={4}>
        <Grid container>
          <Grid xs={12}>
            {data.map((d) => (
              <Box key={d.id} sx={{ minWidth: 275, p: 2 }}>
                <Typography variant="h3">
                  <Link href={getLink(d)}>{d.id}</Link>
                </Typography>
                <Typography variant="p">
                  Status: {d.status}
                </Typography>
                <Grid>
                  <TableContainer component={Paper}>
                    <Table sx={{ minWidth: 650 }} aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell>Name</TableCell>
                          <TableCell>Reps</TableCell>
                          <TableCell>Weight</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {d.exercises.map((e) => (
                          <TableRow
                            key={`${e.name}-${e.set}`}
                            sx={{
                              "&:last-child td, &:last-child th": { border: 0 },
                            }}
                          >
                            <TableCell component="th" scope="row">
                              <ExerciseTitle name={e.name} set={e.set} />
                            </TableCell>
                            <TableCell>{e.reps}</TableCell>
                            <TableCell>{e.weight}</TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Grid>
                <Grid container>
                  <Grid>
                    <TextField
                      sx={{ my: 1, mx: 1 }}
                      label="Session Name"
                      variant="outlined"
                      value={newSessionName}
                      onChange={handleNewSessionNameChange}
                    />
                  </Grid>
                  <Grid>
                    <Button
                      variant="contained"
                      sx={{ my: 2 }}
                      onClick={() => startSessionFromSession(d.id)}
                      item
                      style={{ display: "flex" }}
                    >
                      Start Session
                    </Button>
                  </Grid>
                </Grid>
                <Divider sx={{ my: 2 }} />
              </Box>
            ))}
          </Grid>
        </Grid>
      </Box>
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
            Update Session {session.id}
          </Typography>
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell>Set</TableCell>
                  <TableCell>Name</TableCell>
                  <TableCell>Reps</TableCell>
                  <TableCell>Weight</TableCell>
                  <TableCell></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {session.exercises.map((e) => (
                  <TableRow
                    key={`${e.name}-${e.set}`}
                    sx={{
                      "&:last-child td, &:last-child th": { border: 0 },
                    }}
                  >
                    <TableCell>{e.set}</TableCell>
                    <TableCell component="th" scope="row">
                      {e.name}
                    </TableCell>
                    <TableCell>
                      <TextField
                        sx={{ my: 2 }}
                        label="Reps"
                        variant="outlined"
                        value={e.reps}
                        onChange={(evt) => handleRepsChange(evt, e)}
                      />
                    </TableCell>
                    <TableCell>
                      <TextField
                        sx={{ my: 2 }}
                        label="Weight"
                        variant="outlined"
                        value={e.weight}
                        onChange={(evt) => handleWeightChange(evt, e)}
                      />
                    </TableCell>
                    <TableCell></TableCell>
                  </TableRow>
                ))}
                <TableRow
                  sx={{
                    "&:last-child td, &:last-child th": { border: 0 },
                  }}
                ></TableRow>
              </TableBody>
            </Table>
          </TableContainer>
          <Button sx={{ my: 2 }} variant="outlined" onClick={saveSession}>
            Save Session
          </Button>
        </Box>
      </Modal>
    </div>
  );
}

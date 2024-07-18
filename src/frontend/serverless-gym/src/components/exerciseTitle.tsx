import { useState } from "react";
import {
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Modal,
} from "@mui/material";
import Box from "@mui/material/Box";
import { api } from "../axiosConfig";
import { ExerciseHistroy } from "../models/exerciseHistory";

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

type ExerciseTitleProps = {
  name: string;
  set: number;
};

export default function ExerciseTitle(props: ExerciseTitleProps) {
  const [exerciseHistory, setExerciseHistory] = useState<ExerciseHistroy>({
    name: "",
    history: [],
  });
  const [openModal, setOpenModal] = useState(false);
  const handleModalClose = (
    event: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }

    setExerciseHistory({ name: "", history: [] });
    setOpenModal(false);
  };

  const titleClick = async (exerciseName: string) => {
    const data = await api.get<ExerciseHistroy>(`/history/${exerciseName}`);

    setExerciseHistory(data.data);

    setOpenModal(true);
  };

  return (
    <div>
      <Typography onClick={() => titleClick(props.name)}>
        {props.set}. {props.name}
      </Typography>
      <Modal
        open={openModal}
        onClose={handleModalClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography
            id="modal-modal-title"
            variant="h5"
            component="h5"
            sx={{ my: 2 }}
          >
            {exerciseHistory.name}
          </Typography>
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell>Date</TableCell>
                  <TableCell>Reps x Weight</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {exerciseHistory.history.map((e, index) => (
                  <TableRow
                    key={`${e.date}-${index}`}
                    sx={{
                      "&:last-child td, &:last-child th": { border: 0 },
                    }}
                  >
                    <TableCell component="th" scope="row">
                      {e.date}
                    </TableCell>
                    <TableCell>
                      {e.reps} @ {e.weight}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Modal>
    </div>
  );
}

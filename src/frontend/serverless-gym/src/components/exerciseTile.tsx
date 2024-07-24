import { Card, CardContent, Grid, IconButton, TextField } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import { ApiService } from "../services/apiService";
import { useNavigate } from "react-router-dom";
import { SessionExercise } from "../models/session";
import ExerciseTitle from "./exerciseTitle";

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

type ExerciseTileProps = {
  exercise: SessionExercise;
  readOnly: boolean;
  onDeleteExercise: (e: SessionExercise) => void;
  onRepsChange: (evt: any, e: SessionExercise) => void;
  onWeightChange: (evt: any, e: SessionExercise) => void;
  onWeightBlur: (evt: any) => void;
};

export default function ExerciseTile(props: ExerciseTileProps) {
  const navigate = useNavigate();
  const apiService = new ApiService(navigate);

  return (
    <Card sx={{ margin: "1rem" }}>
      <CardContent>
        <Grid container>
          <Grid xs={10}>
            <ExerciseTitle
              name={props.exercise.name}
              set={props.exercise.set}
            />
          </Grid>
          {props.readOnly ? (
            <div></div>
          ) : (
            <IconButton
              aria-label="delete"
              style={{ float: "right" }}
              onClick={() => props.onDeleteExercise(props.exercise)}
            >
              <DeleteIcon />
            </IconButton>
          )}
          <Grid xs={12}>
            <TextField
              sx={{ my: 2 }}
              label="Reps"
              variant="outlined"
              type="number"
              disabled={props.readOnly}
              value={props.exercise.reps}
              onChange={(evt) => props.onRepsChange(evt, props.exercise)}
            />
          </Grid>
          <Grid xs={12}>
            <TextField
              sx={{ my: 2 }}
              label="Weight"
              variant="outlined"
              type="number"
              disabled={props.readOnly}
              value={props.exercise.weight}
              onChange={(evt) => props.onWeightChange(evt, props.exercise)}
              onBlur={props.onWeightBlur}
            />
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

import { initializeApp } from "firebase/app";
import { useNavigate } from "react-router-dom";
import { firebaseConfig } from "../config";
import {
  getAuth,
  getRedirectResult,
  GoogleAuthProvider,
  signInWithRedirect,
  signOut,
} from "firebase/auth";
import { Box, Button, Grid } from "@mui/material";

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

getRedirectResult(auth)
  .then((result) => {
    if (result === null) {
      return;
    }

    const user = result!.user;

    user.getIdToken(true).then((token) => {
      localStorage.setItem("token", token);
      localStorage.setItem("refreshToken", user.refreshToken);
      localStorage.setItem("emailAddress", user.email!);

      window.location.href = "/";
    });
  })
  .catch((error) => {
    console.log(error);
  });

export default function Login() {
  async function toggleSignIn() {
    if (!auth.currentUser) {
      const provider = new GoogleAuthProvider();
      
      await signInWithRedirect(auth, provider);

      await getRedirectResult(auth);
    } else {
      signOut(auth);
    }
  }

  return (
    <Box component="section" height={"100vh"} width={"100vw"}>
      <Grid container spacing={2}>
        <Grid
          xs={12}
          display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight={"100vh"}
        >
          <Box sx={{ mt: 1 }}>
            <Button variant="contained" onClick={toggleSignIn}>
              Login
            </Button>
          </Box>
        </Grid>
      </Grid>
    </Box>
  );
}

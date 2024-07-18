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
    // This is the signed-in user
    const user = result!.user;
    // This gives you a Facebook Access Token.
    const credential = GoogleAuthProvider.credentialFromResult(result!)!;

    user.getIdToken(true).then((token) => {
      localStorage.setItem("token", token);
      localStorage.setItem("refreshToken", user.refreshToken);
      localStorage.setItem("emailAddress", user.email!);

      window.location.href = '/';
    });
  })
  .catch((error) => {
    // Handle Errors here.
    const errorCode = error.code;
    const errorMessage = error.message;
    // The email of the user's account used.
    const email = error.customData.email;
    // The AuthCredential type that was used.
    const credential = GoogleAuthProvider.credentialFromError(error);
    // ...
  });

export default function Login() {
  const navigate = useNavigate();

  async function toggleSignIn() {
    if (!auth.currentUser) {
      const provider = new GoogleAuthProvider();

      console.log("Redirecting");

      const res = await signInWithRedirect(auth, provider);

      console.log("Handle result");

      // After returning from the redirect when your app initializes you can obtain the result
      const result = await getRedirectResult(auth);
      if (result) {
      }
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

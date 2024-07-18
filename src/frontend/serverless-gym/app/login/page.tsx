"use client";

import { initializeApp } from "firebase/app";
import { firebaseConfig } from "../config";
import {
  getAuth,
  GoogleAuthProvider,
  signInWithPopup,
  signOut,
} from "firebase/auth";
import { useRouter } from "next/navigation";
import { Button } from "@mui/material";

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

export default function Page() {
  const router = useRouter();

  function toggleSignIn() {
    if (!auth.currentUser) {
      const provider = new GoogleAuthProvider();
      signInWithPopup(auth, provider)
        .then(function (result) {
          if (!result) return;
          const credential = GoogleAuthProvider.credentialFromResult(result);
          // This gives you a Google Access Token. You can use it to access the Google API.
          const token = credential?.accessToken;
          // The signed-in user info.
          const user = result.user;
          const authToken = token ?? "";

          localStorage.setItem("token", authToken);
          localStorage.setItem("refreshToken", user.refreshToken);
          localStorage.setItem("emailAddress", user.email!);

          router.push("/");
        })
        .catch(function (error) {
          // Handle Errors here.
          const errorCode = error.code;
          const errorMessage = error.message;
          // The email of the user's account used.
          const email = error.email;
          // The firebase.auth.AuthCredential type that was used.
          const credential = error.credential;
          if (errorCode === "auth/account-exists-with-different-credential") {
            alert(
              "You have already signed up with a different auth provider for that email."
            );
            // If you are using multiple auth providers on your app you should handle linking
            // the user's accounts here.
          } else {
            console.error(error);
          }
        });
    } else {
      signOut(auth);
    }
  }

  return (
    <div>
      <Button variant="outlined" onClick={toggleSignIn}>
        Sign In
      </Button>
    </div>
  );
}

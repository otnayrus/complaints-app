import { Typography, Grid, TextField, Button } from "@mui/material";
import axios from "axios";
import React, { useEffect, useState } from "react";
import MainLayout from "../components/MainLayout";
import GetJWTPayloadFromToken from "../utils/jwt";

const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState(undefined);

  useEffect(() => {
    const raw = localStorage.getItem("token");
    if (raw) {
      setToken(GetJWTPayloadFromToken(raw));
    }
  }, []);

  const submitForm = (e) => {
    e.preventDefault();

    // send request to server with axios
    // if success, redirect to home page
    // if error, show error message

    axios
      .post("http://localhost:8000/users/login", {
        email: email,
        password: password,
      })
      .then((response) => {
        console.log(response);
        localStorage.setItem("token", response.data.token);
        window.location.href = "/";
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const logout = (e) => {
    e.preventDefault();
    localStorage.removeItem("token");
    alert("Logged out successfully")
    window.location.href = "/login";
  }

  return (
    <MainLayout>
      {token ? (
        <div>
          <Typography variant="h6" component="h6" gutterBottom>
            Logged in as {token.name}
          </Typography>

          <Button onClick={logout}>Log out</Button>
        </div>
      ) : (
        <div>
          <Typography variant="h6" component="h6" gutterBottom>
            Login page
          </Typography>

          <form>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  label="Email"
                  fullWidth
                  variant="outlined"
                  onChange={(e) => setEmail(e.target.value)}
                  autoComplete="off"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  label="Password"
                  type={"password"}
                  fullWidth
                  variant="outlined"
                  onChange={(e) => setPassword(e.target.value)}
                  autoComplete="off"
                />
              </Grid>

              <Grid item xs marginY={2}>
                <Button
                  variant="contained"
                  color="primary"
                  type="submit"
                  onClick={submitForm}
                >
                  Login
                </Button>
              </Grid>
            </Grid>
          </form>
        </div>
      )}
    </MainLayout>
  );
};

export default LoginPage;

import { Typography, Grid, TextField, Button } from "@mui/material";
import axios from "axios";
import React, { useState } from "react";
import MainLayout from "../components/MainLayout";

const CreateUserPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const submitForm = (e) => {
    e.preventDefault();

    // send request to server with axios
    // if success, redirect to home page
    // if error, show error message

    if (password !== confirmPassword) {
      alert("Password and confirm password do not match");
      return;
    }

    axios
      .post("http://localhost:8000/users", {
        email: email,
        password: password,
        name: name,
      })
      .then((response) => {
        console.log(response);
        alert("User created successfully");
        window.location.href = "/login";
      })
      .catch((error) => {
        alert("Error creating user" + error.response.data.error);
        console.log(error);
      });
  };

  return (
    <MainLayout>
      <Typography variant="h6" component="h6" gutterBottom>
        Register page
      </Typography>

      <form>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <TextField
              label="Name"
              fullWidth
              variant="outlined"
              onChange={(e) => setName(e.target.value)}
              autoComplete="off"
            />
          </Grid>
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
          <Grid item xs={12}>
            <TextField
              label="Confirm Password"
              type={"password"}
              fullWidth
              variant="outlined"
              autoComplete="off"
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
          </Grid>

          <Grid item xs marginY={2}>
            <Button
              variant="contained"
              color="primary"
              type="submit"
              onClick={submitForm}
            >
              Register
            </Button>
          </Grid>
        </Grid>
      </form>
    </MainLayout>
  );
};

export default CreateUserPage;

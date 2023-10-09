import { Typography } from "@mui/material";
import React from "react";
import MainLayout from "../components/MainLayout";

const Home = () => {
  return (
    <MainLayout>
        <Typography variant="h6" component="h6" gutterBottom > 
            Home
        </Typography>
    </MainLayout>
  );
};

export default Home;

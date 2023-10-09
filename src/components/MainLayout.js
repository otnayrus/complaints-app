import {
  Grid,
  Container,
  Box,
  ListItemButton,
  ListItemText,
  List,
  ListItem,
  Divider,
} from "@mui/material";
import React from "react";
import styled from "@mui/system/styled";

const Side = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  border: "1px solid",
  borderColor: theme.palette.mode === "dark" ? "#444d58" : "#ced7e0",
  padding: theme.spacing(1),
  borderRadius: "4px",
  textAlign: "center",
  minHeight: "100vh",
}));

const Content = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  borderColor: theme.palette.mode === "dark" ? "#444d58" : "#ced7e0",
  padding: theme.spacing(2),
  borderRadius: "4px",
}));

const MainLayout = ({ children }) => {
  return (
    <Container>
      <Box sx={{ my: 4 }}>
        <Grid container spacing={3}>
          <Grid xs>
            <Side>
              <nav>
                <List>
                  <ListItem disablePadding>
                    <ListItemButton component="a" href="/login">
                      <ListItemText primary="Login" />
                    </ListItemButton>
                  </ListItem>
                  <ListItem disablePadding>
                    <ListItemButton component="a" href="/user/create">
                      <ListItemText primary="Register User" />
                    </ListItemButton>
                  </ListItem>
                  <ListItem disablePadding>
                    <ListItemButton component="a" href="/user/profile">
                      <ListItemText primary="User profile" />
                    </ListItemButton>
                  </ListItem>
                </List>
              </nav>
              <Divider />
              <nav>
                <List>
                  <ListItem disablePadding>
                    <ListItemButton component="a" href="/category">
                      <ListItemText primary="Category List" />
                    </ListItemButton>
                  </ListItem>
                </List>
              </nav>
              <Divider />
              <nav>
                <List>
                  <ListItem disablePadding>
                    <ListItemButton component="a" href="/complaint/create">
                      <ListItemText primary="Submit Complaint" />
                    </ListItemButton>
                  </ListItem>
                  <ListItem disablePadding>
                    <ListItemButton component="a" href="/complaint">
                      <ListItemText primary="Complaint List" />
                    </ListItemButton>
                  </ListItem>
                </List>
              </nav>
            </Side>
          </Grid>
          <Grid xs={9}>
            <Content>{children}</Content>
          </Grid>
        </Grid>
      </Box>
    </Container>
  );
};

export default MainLayout;

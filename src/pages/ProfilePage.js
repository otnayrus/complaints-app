import { Typography, Button, Chip } from "@mui/material";
import React, { useEffect, useState } from "react";
import MainLayout from "../components/MainLayout";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import axios from "axios";
import { Link } from "react-router-dom";
import { StyledTableCell } from "../components/StyledTableCell";
import { StyledTableRow } from "../components/StyledTableRow";

const ProfilePage = () => {
  const [data, setData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState(false);

  useEffect(() => {
    var token = localStorage.getItem("token");

    // Make an API call to fetch data with axios with token
    axios
      .get("http://localhost:8000/users/complaints", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
        console.log(response);
        setIsLoading(false);
        console.log(response.data.complaints, "response.data.complaints");
        setData(response.data.complaints);
      })
      .catch((error) => {
        console.log(error);
        setIsLoading(false);
        setIsError(true);
      });
  }, []);

  return (
    <MainLayout>
      <Typography variant="h6" component="h6" gutterBottom>
        Profile page
      </Typography>
      <Typography variant="subtitle1" component="h6" gutterBottom>
        My complaints
      </Typography>
      <Button
        style={{ marginTop: "1rem", marginBottom: "1rem" }}
        href={"/complaint/create"}
      >
        Create new complaint
      </Button>

      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 700 }}>
          <TableHead>
            <TableRow>
              <StyledTableCell>ID</StyledTableCell>
              <StyledTableCell>Description</StyledTableCell>
              <StyledTableCell>Status</StyledTableCell>
              <StyledTableCell>Action</StyledTableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {isLoading ? (
              <TableRow>
                <TableCell colSpan={6}>Loading...</TableCell>
              </TableRow>
            ) : isError ? (
              <TableRow>
                <TableCell colSpan={6}>
                  Error loading data. Please try again.
                </TableCell>
              </TableRow>
            ) : (
              data.map((row) => (
                <StyledTableRow key={row.name}>
                  <TableCell component="th" scope="row">
                    {row.id}
                  </TableCell>
                  <TableCell>{row.description}</TableCell>
                  <TableCell>
                    {row.status === 2 ? (
                        <Chip label="resolved" color="success"/>
                    ) : (
                        <Chip label="pending" color="warning"/>
                    )}
                  </TableCell>
                  <TableCell>
                    <Button
                      component={Link}
                      to={`/complaint/${row.id}`}
                      variant="outlined"
                      color="primary"
                    >
                      Take Action
                    </Button>
                  </TableCell>
                </StyledTableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
    </MainLayout>
  );
};

export default ProfilePage;

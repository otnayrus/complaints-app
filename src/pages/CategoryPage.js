import { Typography, Button } from "@mui/material";
import React, { useEffect, useState } from "react";
import MainLayout from "../components/MainLayout";
import { styled } from "@mui/material/styles";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell, { tableCellClasses } from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import axios from "axios";
import { Link } from "react-router-dom";

const CategoryPage = () => {
  const [data, setData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isError, setIsError] = useState(false);

  useEffect(() => {
    var token = localStorage.getItem("token");

    // Make an API call to fetch data with axios with token
    axios
      .get("http://localhost:8000/categories", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
        console.log(response);
        setIsLoading(false);
        console.log(response.data.categories, "response.data.categories");
        setData(response.data.categories);
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
        Category Page
      </Typography>
      <Button
        style={{ marginTop: "1rem", marginBottom: "1rem" }}
        href={"/category/create"}
      >
        Create new category
      </Button>

      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 700 }}>
          <TableHead>
            <TableRow>
              <StyledTableCell>ID</StyledTableCell>
              <StyledTableCell>Name</StyledTableCell>
              <StyledTableCell>Extra Fields</StyledTableCell>
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
                  <TableCell>{row.name}</TableCell>
                  <TableCell>
                    {row.extra_fields_schema.map((field) => (
                      <div>
                        {field.name} - {field.field_type}
                      </div>
                    ))}
                  </TableCell>
                  <TableCell>
                    <Button
                      component={Link}
                      to={`/category/${row.id}/update`}
                      variant="outlined"
                      color="primary"
                    >
                      Update
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

export default CategoryPage;

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 14,
  },
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  "&:nth-of-type(odd)": {
    backgroundColor: theme.palette.action.hover,
  },
  // hide last border
  "&:last-child td, &:last-child th": {
    border: 0,
  },
}));

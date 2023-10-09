import {
  Checkbox,
  Divider,
  Grid,
  TextField,
  Typography,
  Box,
  Button,
} from "@mui/material";
import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import MainLayout from "../components/MainLayout";

const ComplaintDetailPage = () => {
  const { id } = useParams();
  const [complaint, setComplaint] = useState({});
  const token = localStorage.getItem("token");

  const [isResolved, setIsResolved] = useState(false);
  const [remarks, setRemarks] = useState("");

  const handleCheckboxChange = (event) => {
    setIsResolved(event.target.checked);
  };

  useEffect(() => {
    axios
      .get(`http://localhost:8000/complaints/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
        console.log(response);
        setComplaint(response.data.complaint);
        setRemarks(response.data.complaint.remarks);
        setIsResolved(response.data.complaint.status === 2);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  const submitForm = (e) => {
    e.preventDefault();
    console.log(isResolved, remarks);
  };

  return (
    <MainLayout>
      <Typography variant="h6" component="h6" gutterBottom>
        Complaint detail page
      </Typography>

      <Grid container spacing={2} sx={{ marginY: 2 }}>
        <Grid container spacing={2} paddingX={2} paddingY={1}>
          <Grid item xs={3}>
            Description
          </Grid>
          <Grid item xs>
            {complaint.description}
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Typography sx={{ marginY: 2 }}>Extra Fields</Typography>
        </Grid>
        <Grid container spacing={2} paddingX={2} paddingY={1}>
          {complaint.extra_fields?.map((field, index) => (
            <Grid container spacing={2} key={index} paddingX={2} paddingY={1}>
              <Grid item xs={3}>
                {field.name}
              </Grid>
              <Grid item xs={9}>
                {field.value}
              </Grid>
            </Grid>
          ))}
        </Grid>
      </Grid>

      <Divider />
      <form>
        <Typography sx={{ marginY: 2 }}>Filled by admin</Typography>
        <TextField
          label="Remarks"
          fullWidth
          variant="outlined"
          value={remarks}
          onChange={(e) => setRemarks(e.target.value)}
          autoComplete="off"
        />
        <Box>
          <Checkbox
            checked={isResolved}
            onChange={handleCheckboxChange}
            color="primary" // Change the color if desired
          />
          <label>Resolved</label>
        </Box>

        <Grid item xs marginY={2}>
          <Button
            variant="contained"
            color="primary"
            type="submit"
            onClick={submitForm}
          >
            Update data
          </Button>
        </Grid>
      </form>
    </MainLayout>
  );
};

export default ComplaintDetailPage;

import {
  Typography,
  TextField,
  Button,
  Grid,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import MainLayout from "../components/MainLayout";
import axios from "axios";
import { useParams } from "react-router-dom";

const UpdateCategoryPage = () => {
  const { id } = useParams();
  const [categoryName, setCategoryName] = useState("");
  const [extraFields, setExtraFields] = useState([]);
  const token = localStorage.getItem("token");

  useEffect(() => {
    axios
      .get(`http://localhost:8000/categories/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((response) => {
        console.log(response);
        setCategoryName(response.data.category.name);
        setExtraFields(response.data.category.extra_fields_schema);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  const handleAddField = () => {
    setExtraFields([...extraFields, { name: "", field_type: "" }]);
  };

  const handleRemoveField = (index) => {
    const updatedFields = [...extraFields];
    updatedFields.splice(index, 1);
    setExtraFields(updatedFields);
  };

  const submitForm = (e) => {
    e.preventDefault();
    console.log(categoryName, extraFields);

    // make json payload
    const jsonPayload = {
      id: Number(id),
      name: categoryName,
      extra_fields_schema: extraFields,
    };

    // send request to server with axios with token
    // if success, redirect to category page
    // if error, show error message
    axios
      .patch("http://localhost:8000/categories", jsonPayload, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      })
      .then((response) => {
        console.log(response);
        alert("Category updated successfully");
        // go to category page
        window.location.href = `/category`
      })
      .catch((error) => {
        console.log(error);
        alert("Error creating category: " + error.response.data.error);
      });
  };

  return (
    <MainLayout>
      <Typography variant="h6" component="h6" gutterBottom>
        Category creation page
      </Typography>

      <Typography variant="subtitle1" gutterBottom>
        Category name
      </Typography>

      <form>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <TextField
              label="Name"
              fullWidth
              variant="outlined"
              value={categoryName}
              onChange={(e) => setCategoryName(e.target.value)}
              autoComplete="off"
            />
          </Grid>

          <Grid item xs={12}>
            <Typography variant="subtitle1" gutterBottom>
              Extra Fields
            </Typography>
          </Grid>

          {extraFields.map((field, index) => (
            <Grid container spacing={2} key={index} paddingX={2} paddingY={1}>
              <Grid item xs>
                <TextField
                  label="Field Name"
                  fullWidth
                  variant="outlined"
                  value={field.name}
                  autoComplete="off"
                  onChange={(e) => {
                    const updatedFields = [...extraFields];
                    updatedFields[index].name = e.target.value;
                    setExtraFields(updatedFields);
                  }}
                />
              </Grid>
              <Grid item xs>
                <FormControl fullWidth variant="outlined">
                  <InputLabel>Type</InputLabel>
                  <Select
                    value={field.field_type}
                    onChange={(e) => {
                      const updatedFields = [...extraFields];
                      updatedFields[index].field_type = e.target.value;
                      setExtraFields(updatedFields);
                    }}
                    label="Type"
                  >
                    <MenuItem value="single_file_image">
                      Single File Image
                    </MenuItem>
                    <MenuItem value="multiple_file_image">
                      Multiple File Image
                    </MenuItem>
                    <MenuItem value="dropdown_selection">
                      Dropdown Selection
                    </MenuItem>
                    <MenuItem value="text">Text</MenuItem>
                    <MenuItem value="text_area">Text Area</MenuItem>
                    <MenuItem value="number">Number</MenuItem>
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={2}>
                <Button
                  onClick={() => handleRemoveField(index)}
                  color="secondary"
                >
                  Delete
                </Button>
              </Grid>
            </Grid>
          ))}

          <Grid item xs={12}>
            <Button
              variant="outlined"
              color="primary"
              onClick={handleAddField}
              fullWidth
            >
              Add Extra Field
            </Button>
          </Grid>

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
        </Grid>
      </form>
    </MainLayout>
  );
};

export default UpdateCategoryPage;

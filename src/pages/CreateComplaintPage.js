import {
  Grid,
  TextField,
  Typography,
  Button,
  MenuItem,
  FormControl,
  InputLabel,
  Select,
} from "@mui/material";
import axios from "axios";
import React, { useEffect, useState } from "react";
import MainLayout from "../components/MainLayout";
import GetJWTPayloadFromToken from "../utils/jwt";

const CreateComplaintPage = () => {
  const [complaint, setComplaint] = useState({});
  const token = localStorage.getItem("token");

  const [categories, setCategories] = useState([]);
  const [activeExtraFields, setActiveExtraFields] = useState([]);
  useEffect(() => {
    axios.get("http://localhost:8000/categories").then((response) => {
      console.log(response);
      setCategories(response.data.categories);
    });
  }, []);

  const [extraFields, setExtraFields] = useState([]);

  const submitForm = (e) => {
    var user = GetJWTPayloadFromToken(token);
    e.preventDefault();
    var payload = {
      description: complaint.description,
      category_id: complaint.category_id,
      extra_fields: extraFields,
      created_by: user.id,
    };

    axios.post("http://localhost:8000/complaints", payload, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }).then((response) => {
      console.log(response);
      alert("Complaint created successfully");
      window.location.href = "/complaint";
    }).catch((error) => {
      console.log(error);
      alert("Error creating complaint: " + error.response.data.error);
    }
    );
  };

  return (
    <MainLayout>
      <Typography variant="h6" component="h6" gutterBottom>
        Create complaint page
      </Typography>

      <Grid container spacing={2} sx={{ marginY: 2 }}>
        <Grid
          container
          spacing={2}
          paddingX={2}
          paddingY={1}
          alignItems="center"
        >
          <Grid item xs={3}>
            Description
          </Grid>
          <Grid item xs>
            <TextField
              label="Complaint description"
              fullWidth
              variant="outlined"
              value={complaint.description}
              onChange={(e) =>
                setComplaint({ ...complaint, description: e.target.value })
              }
              autoComplete="off"
            />
          </Grid>
        </Grid>

        <Grid
          container
          spacing={2}
          paddingX={2}
          paddingY={1}
          alignItems="center"
        >
          <Grid item xs={3}>
            Complaint category
          </Grid>
          <Grid item xs>
            <FormControl fullWidth variant="outlined">
              <InputLabel>Category</InputLabel>
              <Select
                value={complaint.category_id}
                onChange={(e) => {
                  setComplaint({ ...complaint, category_id: e.target.value });
                  setActiveExtraFields(
                    categories.find(
                      (category) => category.id === e.target.value
                    ).extra_fields_schema
                  );
                  // init empty array for extra fields
                  setExtraFields(
                    categories
                      .find((category) => category.id === e.target.value)
                      .extra_fields_schema.map((field) => ({
                        name: field.name,
                        value: undefined,
                        field_type: field.field_type,
                      }))
                  );
                }}
                label="Category"
              >
                {categories.map((category, index) => (
                  <MenuItem value={category.id} key={index}>
                    {category.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Typography sx={{ marginY: 2 }}>Extra Fields</Typography>
        </Grid>
        {complaint.category_id ? (
          activeExtraFields.map((field, index) => (
            <Grid
              container
              spacing={2}
              key={index}
              paddingX={2}
              paddingY={1}
              alignItems="center"
            >
              <Grid item xs={3}>
                {field.name}
              </Grid>
              <Grid item xs>
                <DynamicForm
                  name={field.name}
                  field_type={field.field_type}
                  value={field.value}
                  setValue={(value) => {
                    let newExtraFields = [...extraFields];
                    newExtraFields[index].value = value;
                    setExtraFields(newExtraFields);
                  }}
                />
              </Grid>
            </Grid>
          ))
        ) : (
          <Grid item xs>
            Please select the category
          </Grid>
        )}
      </Grid>

      <Grid item xs marginY={2}>
        <Button
          variant="contained"
          color="primary"
          type="submit"
          onClick={submitForm}
        >
          Create complaint
        </Button>
      </Grid>
    </MainLayout>
  );
};

export default CreateComplaintPage;

const DynamicForm = ({ name, field_type, value, setValue }) => {
  switch (field_type) {
    case "text":
      return (
        <TextField
          label="Text"
          fullWidth
          variant="outlined"
          value={value}
          autoComplete="off"
          onChange={(e) => {
            setValue(e.target.value);
          }}
        />
      );
    case "number":
      return (
        <TextField
          type="number"
          label="Number"
          fullWidth
          variant="outlined"
          autoComplete="off"
          onChange={(e) => {
            setValue(e.target.value);
          }}
        />
      );
    case "textarea":
      return (
        <TextField
          label="Textarea"
          fullWidth
          variant="outlined"
          multiline
          rows={4}
          autoComplete="off"
          onChange={(e) => {
            setValue(e.target.value);
          }}
        />
      );
    default:
      return null;
  }
};

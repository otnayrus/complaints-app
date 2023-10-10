import {
  Grid,
  TextField,
  Typography,
  Button,
  MenuItem,
  FormControl,
  InputLabel,
  Select,
  Input,
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

    // Initialize fields and promises arrays
    var fields = [];
    var promises = [];

    extraFields.forEach((el) => {
      if (
        el.field_type === "single_file_image" ||
        el.field_type === "multiple_file_images"
      ) {
        var files = Array.from(el.value);
        var uploaded = [];

        files.forEach((file) => {
          const formData = new FormData();
          formData.append("image", file);

          // Create an Axios request and push it to the promises array
          promises.push(
            axios
              .post("http://localhost:8000/images", formData, {
                headers: {
                  Authorization: `Bearer ${token}`,
                  "Content-Type": "multipart/form-data",
                },
              })
              .then((response) => {
                // Handle the response from the server if needed
                console.log(response.data);
                uploaded.push(response.data.path); // Collect uploaded paths
              })
              .catch((error) => {
                // Handle any errors that occur during the Axios request
                console.error("Error uploading image:", error);
                alert("Error uploading image: " + error);
              })
          );
        });

        // After processing all files for this field, update the fields array
        fields.push({
          field_type: el.field_type,
          name: el.name,
          value: uploaded, // Use the collected uploaded paths
        });
      } else {
        fields.push(el);
      }
    });

    // Use Promise.all to wait for all Axios requests to complete
    Promise.all(promises)
      .then(() => {
        // Rest of your code to create the payload and make the final POST request
        var payload = {
          description: complaint.description,
          category_id: complaint.category_id,
          extra_fields: fields,
          created_by: user.user_id,
        };
        console.log("Payload", payload);

        // Continue with your Axios POST request and handling of the response
        axios
          .post("http://localhost:8000/complaints", payload, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          })
          .then((response) => {
            console.log(response);
            alert("Complaint created successfully");
            window.location.href = "/complaint";
          })
          .catch((error) => {
            console.log(error);
            alert("Error creating complaint: " + error.response.data.error);
          });
      })
      .catch((error) => {
        console.error("Error uploading images:", error);
        alert("Error uploading images: " + error);
      });
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
  const [fileNames, setFileNames] = useState([]);

  switch (field_type) {
    case "text":
      return (
        <Grid
          container
          spacing={2}
          paddingX={2}
          paddingY={1}
          alignItems="center"
        >
          <Grid item xs={3}>
            {name}
          </Grid>
          <Grid item xs>
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
          </Grid>
        </Grid>
      );
    case "number":
      return (
        <Grid
          container
          spacing={2}
          paddingX={2}
          paddingY={1}
          alignItems="center"
        >
          <Grid item xs={3}>
            {name}
          </Grid>
          <Grid item xs>
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
          </Grid>
        </Grid>
      );
    case "text_area":
      return (
        <Grid container spacing={2} paddingX={2} paddingY={1}>
          <Grid item xs={3}>
            {name}
          </Grid>
          <Grid item xs>
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
          </Grid>
        </Grid>
      );
    case "single_file_image":
      return (
        <Grid container spacing={2} paddingX={2} paddingY={1}>
          <Grid item xs={3}>
            {name}
          </Grid>
          <Grid item xs>
            <input
              style={{ display: "none" }}
              accept="image/*"
              id="file-upload"
              type="file"
              onChange={(e) => {
                setFileNames(
                  Array.from(e.target.files).map((file) => file.name)
                );
                setValue(e.target.files);
              }}
            />
            <label htmlFor="file-upload">
              <Button variant="contained" component="span">
                Upload File
              </Button>
            </label>
            <Grid container spacing={2} marginY={1}>
              {fileNames.map((filename, index) => {
                return (
                  <Grid item xs={12}>
                    <Typography key={index}>{filename}</Typography>
                  </Grid>
                );
              })}
            </Grid>
          </Grid>
        </Grid>
      );
    case "multiple_file_images":
      return (
        <Grid container spacing={2} paddingX={2} paddingY={1}>
          <Grid item xs={3}>
            {name}
          </Grid>
          <Grid item xs>
            <input
              style={{ display: "none" }}
              accept="image/*"
              id="file-uploads"
              type="file"
              multiple
              onChange={(e) => {
                const files = e.target.files;
                if (files.length > 0) {
                  setFileNames(Array.from(files).map((file) => file.name));
                  setValue(files);
                }
              }}
            />
            <label htmlFor="file-uploads">
              <Button variant="contained" component="span">
                Upload Multiple Files
              </Button>
            </label>
            <Grid container spacing={2} marginY={1}>
              {fileNames.map((filename, index) => {
                return (
                  <Grid item xs={12}>
                    <Typography key={index}>{filename}</Typography>
                  </Grid>
                );
              })}
            </Grid>
          </Grid>
        </Grid>
      );
    default:
      return null;
  }
};

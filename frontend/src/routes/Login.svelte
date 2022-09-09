<Paper>
  <Title>LOGIN</Title>
  <Content>
    <Textfield 
      bind:value={email}
      type="email"
      bind:invalid={invalid_email}
      updateInvalid
    >
      <svelte:fragment slot="label">
        <CommonIcon
          class="material-icons"
          style="font-size: 1em; line-height: normal; vertical-align: top;"
          >email</CommonIcon
        > Email
      </svelte:fragment>
      <svelte:fragment slot="helper">
        <HelperText></HelperText>
      </svelte:fragment>
    </Textfield>

    <Textfield
      bind:value={password}
      type="password"
      input$maxlength={40}
      bind:invalid={invalid_password}
      updateInvalid
    >
      <svelte:fragment slot="label">
        <CommonIcon
          class="material-icons"
          style="font-size: 1em; line-height: normal; vertical-align: top;"
          >password</CommonIcon
        > Password
      </svelte:fragment>
      <svelte:fragment slot="helper">
        <HelperText>test</HelperText>
      </svelte:fragment>
    </Textfield>

    <Button 
      on:click={login}
      variant="raised"
      disabled={
        invalid_email 
        || invalid_password 
        || email === '' 
        || password === ''
      }
    >
      <Label>LOGIN</Label>
    </Button>

    <Snackbar bind:this={snackbarError} class="password-error">
      <Label>Wrong password and/or wrong username.</Label>
      <Actions>
        <IconButton class="material-icons" title="Dismiss">close</IconButton>
      </Actions>
    </Snackbar>
  </Content>
</Paper>

<script>
  import postData from '../lib/api.js'
  import {push} from 'svelte-spa-router'

  import Paper, { Title, Subtitle, Content } from '@smui/paper';
  import { Icon as CommonIcon } from '@smui/common';

  import Textfield from '@smui/textfield';
  import HelperText from '@smui/textfield/helper-text';
  import Button, { Label } from '@smui/button';

  import SnackbarComponentDev from '@smui/snackbar';
  import Snackbar, { Actions } from '@smui/snackbar';
  import IconButton from '@smui/icon-button';

  let invalid_email = true; let invalid_password = true;
  let snackbarError;
  let email = ''; let password = '';

  function login() {
    let base_url = import.meta.env.VITE_API_URL;
    postData("http://" + base_url + "/login", {
      "username": email,
      "password": password
    })
    .then((response) => {
      if (response.status === 200) {
        console.log("connection successful");
        push("/app");
      } else {
        console.log("connection failed");
        snackbarError && snackbarError.open();
      }
    });
  }

</script>

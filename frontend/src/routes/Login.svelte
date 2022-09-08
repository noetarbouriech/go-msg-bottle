<Paper>
  <Title>LOGIN</Title>
  <Content>
    <Textfield bind:value={email} type="email">
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

    <Textfield bind:value={password} type="password" input$maxlength={40}>
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

    <Button on:click={login} variant="raised">
      <Label>LOGIN</Label>
    </Button>
  </Content>
</Paper>

<script>
  import postData from '../lib/api.js'
  import Paper, { Title, Subtitle, Content } from '@smui/paper';
  import { Icon as CommonIcon } from '@smui/common';

  import Textfield from '@smui/textfield';
  import HelperText from '@smui/textfield/helper-text';
  import Button, { Label } from '@smui/button';

  let email = ''; let password = '';

  function login() {
    let base_url = import.meta.env.VITE_API_URL;
    postData("http://" + base_url + "/login", {
      "username": email,
      "password": password
    })
    .then((data) => {
      console.log(data);
    });
  }

</script>

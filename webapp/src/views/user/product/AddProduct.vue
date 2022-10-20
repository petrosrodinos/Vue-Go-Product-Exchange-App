<template>
  <div class="q-pa-md max-width" style="max-width: 800px">
    <q-stepper v-model="step" ref="stepper" color="primary" animated>
      <q-step
        :name="1"
        :title="id ? 'Edit Product' : 'Create Product'"
        :done="step > 1"
      >
        <CreateProduct
          :formRef="productForm"
          :onImageChange="onImageChange"
          :formState="formState"
          :id="id"
        />
      </q-step>

      <q-step
        :beforeNext="onSubmit"
        :name="2"
        title="Exchange for"
        caption="Optional"
        icon="o_change_circle_utlined"
        :done="step > 2"
      >
        <ExchangeFor :formState="formState" />
      </q-step>

      <q-step
        :name="3"
        :title="id ? 'Product saved' : 'Product Created'"
        icon="o_check_utlined"
        :done="step > 3"
      >
        <h3>
          {{
            id ? "Your product has been saved" : "Your product has been created"
          }}
        </h3>
      </q-step>

      <template v-slot:navigation>
        <q-stepper-navigation>
          <q-btn
            :loading="isLoading && step == 2"
            @click="
              step < 2
                ? validateForm($refs)
                : step == 2
                ? onSubmit($refs)
                : goToProducts()
            "
            color="primary"
            :label="
              step === 2
                ? id
                  ? 'Save'
                  : 'Create'
                : step === 3
                ? 'go to your products'
                : 'Next'
            "
          />
          <q-btn
            v-if="step < 3 && step > 1"
            flat
            color="primary"
            @click="$refs.stepper.previous()"
            label="Back"
            class="q-ml-sm"
          />
        </q-stepper-navigation>
      </template>
    </q-stepper>
  </div>
</template>

<script>
import { ref, reactive, computed } from "vue";
import CreateProduct from "../../../components/AddProduct/CreateProduct.vue";
import ExchangeFor from "../../../components/AddProduct/ExchangeFor.vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { displayErrorMessage } from "@/utils/errorMessage";

export default {
  name: "AddProduct",
  props: ["id"],
  components: {
    CreateProduct,
    ExchangeFor,
  },
  setup(props) {
    const productForm = ref();
    const store = useStore();
    const router = useRouter();
    const formState = reactive({
      name: "",
      description: "",
      quantity: "1",
      image: {},
      type: "",
      listed: true,
      exchangeFor: {
        name: "",
        quantity: "1",
      },
    });

    function validateForm(refs) {
      refs.stepper.next();
      // productForm.value.validate().then((valid) => {
      //   if (valid) {
      //     console.log("valid");
      //   } else {
      //     displayErrorMessage("Please fill out all fields");
      //   }
      // });
    }

    function onSubmit(refs) {
      store
        .dispatch("uploadPhoto", formState.image)
        .then((res) => {
          if (res) {
            store
              .dispatch("createProduct", { ...formState, image: res })
              .then((res) => {
                if (res) {
                  refs.stepper.next();
                }
              })
              .catch((err) => {
                displayErrorMessage("Something went wrong,please try again");
              });
          }
        })
        .catch((err) => {
          console.log(err);
          displayErrorMessage("Something went wrong,please try again");
        });
    }

    function onImageChange(image) {
      formState.image = image;
    }

    function goToProducts() {
      router.push("/profile/products");
    }

    return {
      step: ref(1),
      formState,
      onSubmit,
      isLoading: computed(() => store.state.product.loading),
      onImageChange,
      goToProducts,
      validateForm,
      productForm,
    };
  },
};
</script>

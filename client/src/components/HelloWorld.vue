<template>
  <div class="container mx-auto grid justify-items-start">
    <div class="w-full">
      <div class="flex flex-col">
          <sms-row
              v-for="sms in smsList"
              :key="sms.uuid"
              :from="sms.from"
                   :to="sms.to"
                   :text="sms.text"
                   :when="sms.when"
          />
      </div>
    </div>
  </div>
</template>

<script>
import moment from 'moment';
import Api from '@/service/Api';
import SmsRow from '@/components/SmsRow.vue';

export default {
  name: 'HelloWorld',
  components: { SmsRow },
  props: {
    msg: String,
  },
  data() {
    return {
      smsList: [],
    };
  },

  mounted() {
    Api.getSmsList()
      .then((data) => {
        this.smsList = data.Data;
      });
  },
  methods: {
    formatDate(date) {
      return moment(date).fromNow();
    },
  },
};
</script>

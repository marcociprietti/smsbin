<template>
  <div class="flex p-2 w-full border rounded my-2 cursor-pointer">
    <div class="flex flex-grow mr-10 flex-col items-start w-2/3">

      <div class="flex items-center pb-2">
        <div class="text-lg font-bold mr-3">{{ to }}</div>
        <div class="text-sm text-gray-500">sent by: {{ from }}</div>
      </div>

      <div class="text-gray-500 line-clamp-2 text-left whitespace-pre-line">{{ text }}</div>
    </div>
    <div class="flex flex-col items-end">
      <div class="capitalize pb-2">{{ whenText }}</div>
      <div class="text-gray-500 text-sm">{{ date }}</div>
    </div>
  </div>
</template>

<script>
import moment from 'moment';

export default {
  name: 'SmsRow',
  props: {
    from: {
      type: String,
      required: true,
    },
    to: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
    when: {
      type: String,
      required: true,
    },
  },

  computed: {
    whenText() {
      return moment(this.when).fromNow();
    },
    date() {
      return moment(this.when).format('HH:mm:ss DD-MM-YYYY');
    },
  },
};
</script>
